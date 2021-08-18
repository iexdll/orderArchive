package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"orderArchive/models"
	"orderArchive/store/postgreSQL"
)

var (
	configPath2 string
)

func init() {
	flag.StringVar(&configPath2, "config-path", "configs.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := models.NewConfig()
	_, err := toml.DecodeFile(configPath2, config)
	if err != nil {
		log.Fatal(err)
	}

	dbPostgreSQL, err := postgreSQL.Connection(config.DBPostgreSQL)
	if err != nil {
		log.Fatal(err)
	}
	storeP := postgreSQL.New(dbPostgreSQL)

	customer := "8e5d3d3a-84a1-11e1-97d7-0015179b1a7d"
	tradePoint := models.EmptyRef
	goods := []string{"4e9c219e-4408-11dc-b133-000e0c3ed2d6", "5a65c8cf-1215-11de-af36-000e0ce9b9aa"}

	//ordersID, err := storeP.Order().FindIDByCustomer(customer, tradePoint, 10, -1)
	ordersID, err := storeP.Order().FindIDByGoods(customer, tradePoint, goods)
	log.Println(err)
	log.Println(ordersID)

	var o []*models.Order
	for _, id := range ordersID {
		order, err := storeP.Order().Get(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(order)
		o = append(o, order)
	}

}
