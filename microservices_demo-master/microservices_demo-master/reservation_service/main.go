package main

import (
	"resevation/startup"
	cfg "resevation/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
