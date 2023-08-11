package config

import "os"

type Config struct {
	Port                      string
	ReservationDBHost         string
	ReservationDBPort         string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	LeaveRatingCommandSubject string
	LeaveRatingReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                      "8000",
		ReservationDBHost:         "reservation_db",
		ReservationDBPort:         "27017",
		NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
		LeaveRatingCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		LeaveRatingReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}
