package models

const EmptyRef = "00000000-0000-0000-0000-000000000000"

type Config struct {
	BindAddr     string `toml:"bind_addr"`
	LogLevel     string `toml:"log_level"`
	DBPostgreSQL string `toml:"db_postgresql"`
	DBMongo      string `toml:"db_mongo"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8888",
		LogLevel: "debug",
	}
}
