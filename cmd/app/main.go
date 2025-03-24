package main

import (
	"log"

	"github.com/fun-activities/lottery-backend/config"
	"github.com/fun-activities/lottery-backend/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
