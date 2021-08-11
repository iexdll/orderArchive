package models

import "time"

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
