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

	customer := "11a4e868-e826-43a9-9e13-a86f34581ec6"
	tradePoint := models.EmptyRef

	ordersID, err := storeP.Order().FindIDByCustomer(customer, tradePoint, 10, -1)
	//log.Println(err)
	//log.Println(orders)

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
