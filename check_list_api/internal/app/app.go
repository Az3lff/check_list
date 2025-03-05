package app

import (
	"fmt"

	"github.com/Az3lff/check_list/internal/app/server"
	"github.com/Az3lff/check_list/internal/config"
	"github.com/Az3lff/check_list/internal/repository"
	"github.com/Az3lff/check_list/internal/service"
)

func Run(cfg *config.Config) error {
	//TODO: add pgxpool

	//if err != nil {
	//	return fmt.Errorf("failed connecting to PostgreSQL database: %w", err)
	//}

	repos := repository.NewRepositories(nil)
	services := service.NewServices(repos)

	if err := server.New(cfg.GRPCServer, services).Start(); err != nil {
		return fmt.Errorf("failed starting server: %w", err)
	}

	return nil
}
