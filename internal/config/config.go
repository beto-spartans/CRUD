package config

import (
	"fmt"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresHost     string
	PostgresPort     int
	SSLMode          string
}

func NewConfig() *Config {
	return &Config{
		PostgresUser:     "postgres",
		PostgresPassword: "macbuuren12",
		PostgresDBName:   "users",
		PostgresHost:     "localhost",
		PostgresPort:     5432,
		SSLMode:          "disable",
	}
}

func (c *Config) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDBName,
		c.SSLMode,
	)
}
