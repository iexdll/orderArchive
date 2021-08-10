package main

import (
	"flag"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs.toml", "path to config file")
}

func main() {

}

//Создание миграции
//C:\Users\SmirnovA\go\migrate.exe create -ext sql -dir migrations create_orders
//C:\Users\SmirnovA\go\migrate.exe -path migrations -database "postgres://192.168.99.100:5432/orders?sslmode=disable&user=postgres&password=qwerty" up
