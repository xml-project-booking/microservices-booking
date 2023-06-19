package config

import "os"

type Config struct {
	Port              string
	CatalogueHost     string
	CataloguePort     string
	OrderingHost      string
	OrderingPort      string
	ShippingHost      string
	ShippingPort      string
	InventoryHost     string
	InventoryPort     string
	UserHost          string
	UserPort          string
	AccommodationHost string
	AccommodationPort string
	ReservationHost   string
	ReservationPort   string
	TermHost          string
	TermPort          string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		CatalogueHost:     os.Getenv("CATALOGUE_SERVICE_HOST"),
		CataloguePort:     os.Getenv("CATALOGUE_SERVICE_PORT"),
		OrderingHost:      os.Getenv("ORDERING_SERVICE_HOST"),
		OrderingPort:      os.Getenv("ORDERING_SERVICE_PORT"),
		ShippingHost:      os.Getenv("SHIPPING_SERVICE_HOST"),
		ShippingPort:      os.Getenv("SHIPPING_SERVICE_PORT"),
		InventoryHost:     os.Getenv("INVENTORY_SERVICE_HOST"),
		InventoryPort:     os.Getenv("INVENTORY_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		ReservationHost:   os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
		TermHost:          os.Getenv("TERM_SERVICE_HOST"),
		TermPort:          os.Getenv("TERM_SERVICE_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
	}
}
