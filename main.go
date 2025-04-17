package main

import (
	"Region-Simulator/Config"
	"Region-Simulator/internal/api"
	"log"
)

func main() {
	cfg, err := Config.SetupEnv()

	if err != nil {
		log.Fatalf("Config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
