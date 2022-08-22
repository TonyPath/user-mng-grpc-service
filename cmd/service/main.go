package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	// 3rd party
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/config"
	"github.com/TonyPath/user-mng-grpc-service/internal/repo/sql"
	sqlusers "github.com/TonyPath/user-mng-grpc-service/internal/repo/sql/user"
	"github.com/TonyPath/user-mng-grpc-service/internal/service"
	"github.com/TonyPath/user-mng-grpc-service/logger"
	"github.com/TonyPath/user-mng-grpc-service/stream"
	"github.com/TonyPath/user-mng-grpc-service/transport/grpc"
	httpinfra "github.com/TonyPath/user-mng-grpc-service/transport/http/infra"
)

func main() {
	log, err := logger.New("user-mng-svc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = log.Sync()
	}()

	if err := run(log); err != nil {
		log.Error(err)
		_ = log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		return err
	}

	// Database connection
	// -------------------
	pgCfg := sql.Config{
		User:     cfg.DB.Username,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		DBName:   cfg.DB.DBName,
	}
	db, err := sql.NewDB(pgCfg)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Errorw("close db", "error", err)
		}
	}()

	// App Dependencies
	// ----------------
	usersRepo := sqlusers.NewRepository(db, log)

	streamConfig := stream.Config{
		Brokers: strings.Split(cfg.Kafka.ProducerBrokers, ","),
	}
	kafkaClient, err := stream.NewKafkaClient(streamConfig)
	if err != nil {
		return err
	}
	publisher := stream.NewEventPublisher(kafkaClient)
	defer publisher.Close()

	svc := service.NewUserService(usersRepo, publisher)

	//---------------------------
	//
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	//---------------------------
	//
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, gctx := errgroup.WithContext(cctx)

	g.Go(func() error {
		select {
		case sig := <-shutdown:
			cancel()
			return fmt.Errorf("receive term sign %s", sig)
		case <-gctx.Done():
			return gctx.Err()
		}
	})

	infraServer := httpinfra.NewServer(log, db)
	g.Go(func() error {
		return infraServer.Run(gctx)
	})

	grpcServer := grpc.NewServer(log, fmt.Sprintf(":%d", cfg.GRPCPort), svc)
	g.Go(func() error {
		return grpcServer.Run(gctx)
	})

	return g.Wait()
}
