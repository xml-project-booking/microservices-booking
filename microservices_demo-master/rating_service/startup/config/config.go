package config

import "os"

type Config struct {
	Port                      string
	RatingDBHost              string
	RatingDBPort              string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	LeaveRatingCommandSubject string
	LeaveRatingReplySubject   string
	NoificationSubject        string
}

func NewConfig() *Config {
	return &Config{
		Port:                      "8000",
		RatingDBHost:              "rating_db",
		RatingDBPort:              "27017",
		NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
		LeaveRatingCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		LeaveRatingReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
		NoificationSubject:        os.Getenv("NOTIFICATION_SUBJECT"),
	}
}
