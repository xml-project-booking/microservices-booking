package main

import (
	"term_service/startup"
	cfg "term_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
