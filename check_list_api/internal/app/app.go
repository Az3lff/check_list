package app

import (
	"context"
	"fmt"

	"github.com/Az3lff/check_list/internal/app/server"
	"github.com/Az3lff/check_list/internal/config"
	"github.com/Az3lff/check_list/internal/repository"
	"github.com/Az3lff/check_list/internal/service"
)

func Run(cfg *config.Config) error {
	repo, err := repository.NewTaskRepository(context.Background(), cfg.Postgres)
	if err != nil {
		return fmt.Errorf("failed connecting to PostgreSQL database: %w", err)
	}

	srv := server.New(cfg.GRPCServer, service.NewTaskService(repo))
	if err := srv.Start(); err != nil {
		return fmt.Errorf("failed starting server: %w", err)
	}

	return nil
}
