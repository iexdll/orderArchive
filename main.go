package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
	"log"
	"net"
	"orderArchive/api"
	"orderArchive/models"
	"orderArchive/store/postgreSQL"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := models.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("Ошибка получения настроек: ", err)
	}

	db, err := postgreSQL.Connection(config.DBPostgreSQL)
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}
	store := postgreSQL.New(db)

	server := grpc.NewServer()
	instance := &api.OrderServer{Store: store}
	api.RegisterOrderServiceServer(server, instance)

	listener, err := net.Listen("tcp", config.BindAddr)
	if err != nil {
		log.Fatal("Ошибка создания gRPC:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Ошибка запуска gRPC сервера:", err)
	}

}

//Создание миграции
//C:\Users\SmirnovA\go\migrate.exe create -ext sql -dir migrations create_orders
//C:\Users\SmirnovA\go\migrate.exe -path migrations -database "postgres://192.168.99.100:5432/orders?sslmode=disable&user=postgres&password=qwerty" up
