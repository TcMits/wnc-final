package main

import (
	"log"

	"github.com/TcMits/wnc-final/config"
)

func main() {
	// Configuration
	_, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	// app.Run(cfg)
}
