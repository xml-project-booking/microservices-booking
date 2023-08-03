package config

import "os"

type Config struct {
	Port         string
	RatingDBHost string
	RatingDBPort string
	NatsHost     string
	NatsPort     string
	NatsUser     string
	NatsPass     string
}

func NewConfig() *Config {
	return &Config{
		Port:         "8000",
		RatingDBHost: "rating_db",
		RatingDBPort: "27017",
		NatsHost:     os.Getenv("NATS_HOST"),
		NatsPort:     os.Getenv("NATS_PORT"),
		NatsUser:     os.Getenv("NATS_USER"),
		NatsPass:     os.Getenv("NATS_PASS"),
	}
}
