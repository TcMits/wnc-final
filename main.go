package main

import (
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/internal/app"
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
