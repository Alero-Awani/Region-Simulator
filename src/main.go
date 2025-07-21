package main

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/api"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("This config file is not loaded properly %v\n", err)
		fmt.Println("The error message is:", err)
	}

	api.StartServer(cfg)
}
