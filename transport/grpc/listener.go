package grpc

import (
	"context"
	"errors"
	"net"
	"time"

	// 3rd party
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// internal
	pb "github.com/TonyPath/user-mng-grpc-service/proto/services/user"
)

const (
	defaultConnectionTimeout = 30 * time.Second
	shutdownGracePeriod      = 5 * time.Second
)

type Server struct {
	grpcServer *grpc.Server
	addr       string
	logger     *zap.SugaredLogger
}

func NewServer(
	logger *zap.SugaredLogger,
	addr string,
	svc userService) *Server {
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(defaultConnectionTimeout),
	)
	pb.RegisterUserServer(grpcServer, New(logger, svc))

	/*
		Used mostly for testing under development.
		Enable reflection in order perform grpcurl (https://github.com/fullstorydev/grpcurl) calls
		or use postman beta feature of grpc request.
	*/
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
		addr:       addr,
		logger:     logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		s.logger.Infow("startup", "status", "grpc server started", "host", s.addr)
		errCh <- s.listenAndServe()
	}()

	select {
	case <-ctx.Done():
		return s.shutdown()
	case err := <-errCh:
		return err
	}
}

func (s *Server) listenAndServe() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(l)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownGracePeriod)
	defer cancel()

	go func() {
		<-ctx.Done()
		if !errors.Is(ctx.Err(), context.Canceled) {
			s.grpcServer.Stop()
			s.logger.Infow("shutdown", "status", "forcibly stop grpc server")
		}
	}()

	s.grpcServer.GracefulStop()
	s.logger.Infow("shutdown", "status", "gracefully stopped grpc server")

	return nil
}
