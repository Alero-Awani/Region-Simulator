package main

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
