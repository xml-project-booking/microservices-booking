package config

import "os"

type Config struct {
	Port       string
	TermDBHost string
	TermDBPort string
	NatsHost   string
	NatsPort   string
	NatsUser   string
	NatsPass   string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8000",
		TermDBHost: "term_db",
		TermDBPort: "27017",
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsPort:   os.Getenv("NATS_PORT"),
		NatsUser:   os.Getenv("NATS_USER"),
		NatsPass:   os.Getenv("NATS_PASS"),
	}
}
