package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"orderArchive/models"
	"strings"
	"time"
	"unicode/utf8"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Save(o *models.Order) error {
	return nil
}

func (r *OrderRepository) FindIDByCustomer(customer string, tradePoint string, limit int32, skip int32) ([]string, error) {
	return nil, nil
}

func (r *OrderRepository) FindIDByGoods(customer string, tradePoint string, goodsID []string) ([]string, error) {
	return nil, nil
}

func (r *OrderRepository) Get(id string) (*models.Order, error) {

	orderID := strings.ToUpper(id)
	order := &models.Order{}

	var result bson.M
	collection := r.store.db.Collection("orders")
	err := collection.FindOne(context.TODO(), bson.M{"_id": orderID}).Decode(&result)

	if err != nil {
		//Добавить проверку на mongo: no documents in result
		return nil, err
	}

	orderCancel, err := r.getOrderCancel(orderID)
	if err != nil {
		return nil, err
	}

	order.ID = result["_id"].(string)
	order.Number = result["number"].(string)
	order.Date = result["date"].(primitive.DateTime).Time()
	order.ShippingDate = result["shippingDate"].(primitive.DateTime).Time()
	if result["deliveryDate"] != nil {
		order.DeliveryDate = result["deliveryDate"].(primitive.DateTime).Time()
	}
	if result["warehouseShipping"] != nil {
		order.WarehouseShipping = result["warehouseShipping"].(string)
	} else {
		order.WarehouseShipping = "00000000-0000-0000-0000-000000000000"
	}
	order.Customer = result["customer"].(string)
	order.Sum = result["sum"].(float64)
	order.TradePoint = result["tradePoint"].(string)
	order.PaymentType = result["paymentType"].(int32)
	order.DeliveryType = result["deliveryType"].(int32)
	order.TransportType = result["transportType"].(int32)
	order.Status = result["status"].(int32)
	order.Comment = result["comment"].(string)

	if utf8.RuneCountInString(order.Comment) > 255 {
		order.Comment = string([]rune(order.Comment)[:255])
	}

	goodsList := result["goodsList"].(primitive.A)
	for _, item := range goodsList {
		v := item.(primitive.M)
		orderRow := models.OrderRow{
			RowNumber: v["rowNumber"].(int32),
			RowID:     v["rowID"].(string),
			Goods:     v["goods"].(string),
			Group:     v["warehouseGroup"].(string),
			Price:     v["price"].(float64),
			Count:     v["count"].(int32),
			Comment:   v["comment"].(string),
		}
		key := v["rowID"].(string) + v["goods"].(string)
		if c, ok := orderCancel[key]; ok {
			orderRow.CountCancel = c
		}
		order.GoodsList = append(order.GoodsList, orderRow)
	}

	order.AfterGet()
	return order, nil
}

func (r *OrderRepository) getOrderCancel(id string) (map[string]int32, error) {

	collection := r.store.db.Collection("orderCancel")
	cur, err := collection.Find(context.TODO(), bson.M{"orderID": id, "remove": false})
	if err != nil {
		return nil, err
	}

	orderCancel := make(map[string]int32)

	for cur.Next(context.TODO()) {

		var elem bson.M
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		goodsList := elem["goodsList"].(primitive.A)
		for _, item := range goodsList {
			v := item.(primitive.M)
			key := v["rowID"].(string) + v["goods"].(string)
			count := v["count"].(int32)
			if c, ok := orderCancel[key]; ok {
				count = count + c
			}
			orderCancel[key] = count
		}

	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(context.TODO())

	return orderCancel, nil
}

func (r *OrderRepository) FindIDByDate(dateStart time.Time, dateEnd time.Time) ([]string, error) {

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"_id": 1})
	findOptions.SetSort(bson.M{"date": 1})

	filter := bson.M{
		"date": bson.M{
			"$gte": dateStart, //>=
			"$lte": dateEnd,   ///<=
		},
	}

	collection := r.store.db.Collection("orders")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	result := make([]struct {
		ID string `bson:"_id"`
	}, 0)
	_ = cur.All(context.TODO(), &result)
	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(context.TODO())

	var orders []string
	for _, item := range result {
		orders = append(orders, item.ID)
	}

	return orders, nil
}
