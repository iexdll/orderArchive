package models

type Config struct {
	BindAddr     string `toml:"bind_addr"`
	LogLevel     string `toml:"log_level"`
	DBPostgreSQL string `toml:"db_postgresql"`
	DBMongo      string `toml:"db_mongo"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8089",
		LogLevel: "debug",
	}
}
