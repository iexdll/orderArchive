package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"orderArchive/models"
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

	dbPostgreSQL, err := postgreSQL.Connection(config.DBPostgreSQL)
	if err != nil {
		log.Fatal(err)
	}
	storeP := postgreSQL.New(dbPostgreSQL)

	o, err := storeP.Order().Get("00000f6d-96ce-ae45-7b95-6ce92029fd28")

	log.Println(err)
	log.Println(o)

}
