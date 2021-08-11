package models

import (
	"strings"
	"time"
)

type Order struct {
	ID                string
	Number            string
	Date              time.Time
	ShippingDate      time.Time
	DeliveryDate      time.Time
	Customer          string
	Sum               float64
	TradePoint        string
	PaymentType       int32
	DeliveryType      int32
	TransportType     int32
	Status            int32
	Comment           string
	WarehouseShipping string
	GoodsList         []OrderRow
}

type OrderRow struct {
	RowNumber   int32
	RowID       string
	Goods       string
	Group       string
	Price       float64
	Count       int32
	CountCancel int32
	Comment     string
}

func (o *Order) AfterGet() {

	layout := "2006-01-02"
	o.ShippingDate, _ = time.ParseInLocation(layout, o.ShippingDate.Format(layout), time.Local)
	o.DeliveryDate, _ = time.ParseInLocation(layout, o.DeliveryDate.Format(layout), time.Local)

	o.ID = strings.ToUpper(o.ID)
	o.Customer = strings.ToUpper(o.Customer)
	o.TradePoint = strings.ToUpper(o.TradePoint)
	o.WarehouseShipping = strings.ToUpper(o.WarehouseShipping)

	for i, item := range o.GoodsList {
		o.GoodsList[i].RowID = strings.ToUpper(item.RowID)
		o.GoodsList[i].Goods = strings.ToUpper(item.Goods)
		o.GoodsList[i].Group = strings.ToUpper(item.Group)
	}
}
