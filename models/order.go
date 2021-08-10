package models

import "time"

type Order struct {
	ID                string    //uuid
	Number            string    //varchar(10)
	Date              time.Time //timestamp
	ShippingDate      time.Time //date
	DeliveryDate      time.Time //date
	Customer          string    //uuid
	Sum               float64   //money
	TradePoint        string    //uuid
	PaymentType       int32     //smallint
	DeliveryType      int32     //smallint
	TransportType     int32     //smallint
	Status            int32     //smallint
	Comment           string    //varchar(100)
	WarehouseShipping string    //uuid
	GoodsList         []OrderRow
}

type OrderRow struct {
	RowNumber   int32   //smallint
	RowID       string  //uuid
	Goods       string  //uuid
	Group       string  //uuid
	Price       float64 //money
	Count       int32   //smallint
	CountCancel int32   //smallint
	Comment     string  //varchar(100)
}
