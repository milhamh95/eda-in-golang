package monolith

import (
	"context"
	"database/sql"
	"eda-in-golang/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
}

type Module interface {
	Startup(ctx context.Context, monolith Monolith) error
}
