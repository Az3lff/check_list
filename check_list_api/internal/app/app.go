package app

import (
	"context"
	"fmt"

	"github.com/Az3lff/check_list/internal/app/server"
	"github.com/Az3lff/check_list/internal/config"
	"github.com/Az3lff/check_list/internal/repository"
	"github.com/Az3lff/check_list/internal/service"
	"github.com/Az3lff/check_list/pkg/database/postgres"
)

func Run(cfg *config.Config) error {
	pg := postgres.NewPostgres(context.Background(), cfg.Postgres.Connection)
	pgPool, err := pg.Connection()
	if err != nil {
		return fmt.Errorf("failed connecting to PostgreSQL database: %w", err)
	}

	repos := repository.NewRepositories(pgPool)
	services := service.NewServices(repos)

	if err := server.New(cfg.GRPCServer, services).Start(); err != nil {
		return fmt.Errorf("failed starting server: %w", err)
	}

	return nil
}
