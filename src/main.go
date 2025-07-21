package main

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/api"
)

func main() {
	cfg, _ := config.SetupEnv()

	// if err != nil {
	// 	log.Fatalf("config file is not loaded properly %v\n", err)
	// 	fmt.Println(err)
	// }

	api.StartServer(cfg)
}
