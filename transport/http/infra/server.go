package infra

import (
	"context"
	"database/sql"
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	// 3rd party
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

const (
	httpPort            = 4000
	httpReadTimeout     = 30 * time.Second
	httpWriteTimeout    = 60 * time.Second
	shutdownGracePeriod = 5 * time.Second
)

type Server struct {
	httpServer http.Server
	db         *sql.DB
	logger     *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger, db *sql.DB) *Server {
	debugAPI := &Server{
		httpServer: http.Server{
			Addr:         fmt.Sprintf(":%d", httpPort),
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
		},
		db:     db,
		logger: logger,
	}

	debugAPI.addRoutes()

	return debugAPI
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		s.logger.Infow("startup", "status", "http infra server started")
		errCh <- s.httpServer.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return s.shutdown()
	case err := <-errCh:
		return err
	}
}

func (s *Server) addRoutes() {
	mux := chi.NewRouter()

	pprofRegister(mux)

	ch := checkHandler{
		DB: s.db,
	}

	mux.HandleFunc("/readiness", ch.Readiness)
	mux.HandleFunc("/liveness", ch.Liveness)

	s.httpServer.Handler = mux
}

func pprofRegister(mux *chi.Mux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())
}

func (s *Server) shutdown() error {
	tctx, cancel := context.WithTimeout(context.Background(), shutdownGracePeriod)
	defer cancel()

	if err := s.httpServer.Shutdown(tctx); err != nil {
		closeErr := s.httpServer.Close()
		if closeErr != nil {
			return fmt.Errorf("cannot stop server gracefully: %w", err)
		}
	}
	s.logger.Infow("shutdown", "status", "gracefully stopped http infra server")

	return nil
}
