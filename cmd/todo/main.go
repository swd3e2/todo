package main

import (
	"github.com/swd3e2/todo/internal/application/core"
	"log"
)

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
