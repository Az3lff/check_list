package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Az3lff/check_list/internal/app"
	"github.com/Az3lff/check_list/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Cannot load config: %s", err.Error())
	}

	if err := app.Run(cfg); err != nil {
		logrus.Fatalf("Cannot run app: %s", err.Error())
	}
}
