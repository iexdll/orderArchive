package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"orderArchive/models"
	"orderArchive/store/mongoDB"
	"orderArchive/store/postgreSQL"
)

var (
	configPath1 string
)

func init() {
	flag.StringVar(&configPath1, "config-path", "configs.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := models.NewConfig()
	_, err := toml.DecodeFile(configPath1, config)
	if err != nil {
		log.Fatal(err)
	}

	dbMongo, err := mongoDB.Connection(config.DBMongo)
	if err != nil {
		log.Fatal(err)
	}
	storeM := mongoDB.New(dbMongo)

	dbPostgreSQL, err := postgreSQL.Connection(config.DBPostgreSQL)
	if err != nil {
		log.Fatal(err)
	}
	storeP := postgreSQL.New(dbPostgreSQL)

	orderID := "002106A8-20EF-D584-6059-5AA48AF6AE68"
	orderM, err := storeM.Order().Get(orderID)
	orderP, err := storeP.Order().Get(orderID)
	orderM.AfterGet()
	orderP.AfterGet()

	log.Println(orderM)
	log.Println(orderP)
	log.Println(orderM == orderP)
}
