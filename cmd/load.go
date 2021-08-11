package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"orderArchive/models"
	"orderArchive/store/mongoDB"
	"orderArchive/store/postgreSQL"
	"sync"
	"time"
)

var (
	configPath          string
	countDoneOrders     int
	lockCountDoneOrders sync.Mutex
)

const dateStart = "2017-01-01"
const dateEnd = "2017-09-01"
const countChannel = 30
const countOrderChannel = 1000

func init() {
	flag.StringVar(&configPath, "config-path", "configs.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := models.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
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

	MST, _ := time.LoadLocation("Europe/Moscow")
	dStart, _ := time.ParseInLocation("2006-01-02T15:04:05", dateStart+"T00:00:00", MST)
	dEnd, _ := time.ParseInLocation("2006-01-02T15:04:05", dateEnd+"T23:59:59", MST)

	orders, _ := storeM.Order().FindByDate(dStart, dEnd)
	ordersLen := len(orders)

	log.Printf("Заказов с %v по %v: %d\n", dStart, dEnd, ordersLen)

	c := make(chan Slice, countChannel)
	for i := 0; i < countChannel; i++ {
		worker := &Worker{
			id:     i,
			storeM: storeM,
			storeP: storeP,
			orders: &orders,
		}
		go worker.process(c)
	}

	for i := 0; i < ordersLen; i = i + countOrderChannel {
		low := i
		high := i + countOrderChannel
		if high > ordersLen {
			high = ordersLen
		}
		c <- Slice{Low: low, High: high}
		log.Printf("Отправлено в канал заказы с %d по %d\n", low, high)
	}

	time.Sleep(time.Millisecond * 50000000)

}

type Slice struct {
	Low  int
	High int
}

type Worker struct {
	id     int
	storeM *mongoDB.Store
	storeP *postgreSQL.Store
	orders *[]string
}

func (w *Worker) process(c chan Slice) {
	log.Printf("Запущен %d канал\n", w.id)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		for i := v.Low; i < v.High; i++ {
			oID := (*w.orders)[i]
			order, err := w.storeM.Order().Get(oID)
			if err != nil {
				log.Printf("Ошибка получения заказа %s", oID)
				log.Fatal(err)
			}
			err = w.storeP.Order().Save(order)
			if err != nil {
				log.Printf("Ошибка записи заказа %s", oID)
				log.Fatal(err)
			}

			lockCountDoneOrders.Lock()
			countDoneOrders++
			lockCountDoneOrders.Unlock()

			if countDoneOrders%100 == 0 {
				log.Printf("Обработано %d из %d заказов\n", countDoneOrders, len(*w.orders))
			}
		}
	}
	log.Printf("Канал %d закрыт\n", w.id)
}
