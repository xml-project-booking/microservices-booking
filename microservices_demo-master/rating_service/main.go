package main

import (
	"rating_service/startup"
	cfg "rating_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
