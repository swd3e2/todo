package main

import (
	"github.com/swd3e2/todo/internal/application/core"
	"log"
)

// @title Todo App API
// @version 1.0.0
// @description Demo todo app

// @host localhost:8091
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := core.New()

	if err := app.Configure("app"); err != nil {
		return err
	}

	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
