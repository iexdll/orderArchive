// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.6
// source: orders.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentType int32

const (
	PaymentType_CASH     PaymentType = 0
	PaymentType_CASHLESS PaymentType = 1
	PaymentType_CARD     PaymentType = 2
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "CASH",
		1: "CASHLESS",
		2: "CARD",
	}
	PaymentType_value = map[string]int32{
		"CASH":     0,
		"CASHLESS": 1,
		"CARD":     2,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

type DeliveryType int32

const (
	DeliveryType_DELIVERY DeliveryType = 0
	DeliveryType_PICKUP   DeliveryType = 1
)

// Enum value maps for DeliveryType.
var (
	DeliveryType_name = map[int32]string{
		0: "DELIVERY",
		1: "PICKUP",
	}
	DeliveryType_value = map[string]int32{
		"DELIVERY": 0,
		"PICKUP":   1,
	}
)

func (x DeliveryType) Enum() *DeliveryType {
	p := new(DeliveryType)
	*p = x
	return p
}

func (x DeliveryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[1].Descriptor()
}

func (DeliveryType) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[1]
}

func (x DeliveryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryType.Descriptor instead.
func (DeliveryType) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

type TransportType int32

const (
	TransportType_CAR     TransportType = 0
	TransportType_RAILWAY TransportType = 1
	TransportType_AVIA    TransportType = 2
)

// Enum value maps for TransportType.
var (
	TransportType_name = map[int32]string{
		0: "CAR",
		1: "RAILWAY",
		2: "AVIA",
	}
	TransportType_value = map[string]int32{
		"CAR":     0,
		"RAILWAY": 1,
		"AVIA":    2,
	}
)

func (x TransportType) Enum() *TransportType {
	p := new(TransportType)
	*p = x
	return p
}

func (x TransportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[2].Descriptor()
}

func (TransportType) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[2]
}

func (x TransportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportType.Descriptor instead.
func (TransportType) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

type Status int32

const (
	Status_DEFAULT         Status = 0 //Записан
	Status_PROCESSED       Status = 1 //Обработан
	Status_COLLECT         Status = 2 //В сборке
	Status_COLLECTED       Status = 3 //Собран
	Status_SHIPPED         Status = 4 //Отгружен
	Status_CLOSED          Status = 5 //Закрыт
	Status_CANCELED        Status = 6 //Отменен
	Status_SUSPENDED       Status = 7 //Приостановлен
	Status_NOTCOMPLETED    Status = 8 //Не выполнен
	Status_PAYMENTEXPECTED Status = 9 //Ожидается оплата
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "DEFAULT",
		1: "PROCESSED",
		2: "COLLECT",
		3: "COLLECTED",
		4: "SHIPPED",
		5: "CLOSED",
		6: "CANCELED",
		7: "SUSPENDED",
		8: "NOTCOMPLETED",
		9: "PAYMENTEXPECTED",
	}
	Status_value = map[string]int32{
		"DEFAULT":         0,
		"PROCESSED":       1,
		"COLLECT":         2,
		"COLLECTED":       3,
		"SHIPPED":         4,
		"CLOSED":          5,
		"CANCELED":        6,
		"SUSPENDED":       7,
		"NOTCOMPLETED":    8,
		"PAYMENTEXPECTED": 9,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[3].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[3]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number            string        `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Date              string        `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	ShippingDate      string        `protobuf:"bytes,4,opt,name=shippingDate,proto3" json:"shippingDate,omitempty"`
	DeliveryDate      string        `protobuf:"bytes,5,opt,name=deliveryDate,proto3" json:"deliveryDate,omitempty"`
	Customer          string        `protobuf:"bytes,6,opt,name=customer,proto3" json:"customer,omitempty"`
	Sum               float32       `protobuf:"fixed32,7,opt,name=sum,proto3" json:"sum,omitempty"`
	TradePoint        string        `protobuf:"bytes,8,opt,name=tradePoint,proto3" json:"tradePoint,omitempty"`
	PaymentType       PaymentType   `protobuf:"varint,9,opt,name=paymentType,proto3,enum=orderservice.PaymentType" json:"paymentType,omitempty"`
	DeliveryType      DeliveryType  `protobuf:"varint,10,opt,name=deliveryType,proto3,enum=orderservice.DeliveryType" json:"deliveryType,omitempty"`
	TransportType     TransportType `protobuf:"varint,11,opt,name=transportType,proto3,enum=orderservice.TransportType" json:"transportType,omitempty"`
	Status            Status        `protobuf:"varint,12,opt,name=status,proto3,enum=orderservice.Status" json:"status,omitempty"`
	Comment           string        `protobuf:"bytes,13,opt,name=comment,proto3" json:"comment,omitempty"`
	WarehouseShipping string        `protobuf:"bytes,14,opt,name=warehouseShipping,proto3" json:"warehouseShipping,omitempty"`
	GoodsList         []*OrderRow   `protobuf:"bytes,15,rep,name=GoodsList,proto3" json:"GoodsList,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetShippingDate() string {
	if x != nil {
		return x.ShippingDate
	}
	return ""
}

func (x *Order) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *Order) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *Order) GetSum() float32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Order) GetTradePoint() string {
	if x != nil {
		return x.TradePoint
	}
	return ""
}

func (x *Order) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_CASH
}

func (x *Order) GetDeliveryType() DeliveryType {
	if x != nil {
		return x.DeliveryType
	}
	return DeliveryType_DELIVERY
}

func (x *Order) GetTransportType() TransportType {
	if x != nil {
		return x.TransportType
	}
	return TransportType_CAR
}

func (x *Order) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_DEFAULT
}

func (x *Order) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Order) GetWarehouseShipping() string {
	if x != nil {
		return x.WarehouseShipping
	}
	return ""
}

func (x *Order) GetGoodsList() []*OrderRow {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

type OrderRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowNumber   int32   `protobuf:"varint,1,opt,name=RowNumber,proto3" json:"RowNumber,omitempty"`
	RowID       string  `protobuf:"bytes,2,opt,name=RowID,proto3" json:"RowID,omitempty"`
	Goods       string  `protobuf:"bytes,3,opt,name=Goods,proto3" json:"Goods,omitempty"`
	Group       string  `protobuf:"bytes,4,opt,name=Group,proto3" json:"Group,omitempty"`
	Price       float32 `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	Count       int32   `protobuf:"varint,6,opt,name=Count,proto3" json:"Count,omitempty"`
	CountCancel int32   `protobuf:"varint,7,opt,name=CountCancel,proto3" json:"CountCancel,omitempty"`
	Comment     string  `protobuf:"bytes,8,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *OrderRow) Reset() {
	*x = OrderRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRow) ProtoMessage() {}

func (x *OrderRow) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRow.ProtoReflect.Descriptor instead.
func (*OrderRow) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

func (x *OrderRow) GetRowNumber() int32 {
	if x != nil {
		return x.RowNumber
	}
	return 0
}

func (x *OrderRow) GetRowID() string {
	if x != nil {
		return x.RowID
	}
	return ""
}

func (x *OrderRow) GetGoods() string {
	if x != nil {
		return x.Goods
	}
	return ""
}

func (x *OrderRow) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *OrderRow) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderRow) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrderRow) GetCountCancel() int32 {
	if x != nil {
		return x.CountCancel
	}
	return 0
}

func (x *OrderRow) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0xc5, 0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x75, 0x6d,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x6f,
	0x77, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xd2, 0x01, 0x0a,
	0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x77,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x6f,
	0x77, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x77, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x77, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2a, 0x2f, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41,
	0x53, 0x48, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x52, 0x44,
	0x10, 0x02, 0x2a, 0x28, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x0d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x41, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x41, 0x49, 0x4c, 0x57, 0x41,
	0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x56, 0x49, 0x41, 0x10, 0x02, 0x2a, 0x9d, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x09, 0x32, 0x4c, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x04, 0x2f,
	0x61, 0x70, 0x69, 0xca, 0x02, 0x13, 0x41, 0x70, 0x70, 0x5c, 0x41, 0x50, 0x49, 0x5c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x41, 0x72, 0x68, 0x69, 0x76, 0x65, 0xe2, 0x02, 0x13, 0x41, 0x70, 0x70, 0x5c,
	0x41, 0x50, 0x49, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orders_proto_goTypes = []interface{}{
	(PaymentType)(0),    // 0: orderservice.PaymentType
	(DeliveryType)(0),   // 1: orderservice.DeliveryType
	(TransportType)(0),  // 2: orderservice.TransportType
	(Status)(0),         // 3: orderservice.Status
	(*GetRequest)(nil),  // 4: orderservice.GetRequest
	(*GetResponse)(nil), // 5: orderservice.GetResponse
	(*Order)(nil),       // 6: orderservice.Order
	(*OrderRow)(nil),    // 7: orderservice.OrderRow
}
var file_orders_proto_depIdxs = []int32{
	6, // 0: orderservice.GetResponse.order:type_name -> orderservice.Order
	0, // 1: orderservice.Order.paymentType:type_name -> orderservice.PaymentType
	1, // 2: orderservice.Order.deliveryType:type_name -> orderservice.DeliveryType
	2, // 3: orderservice.Order.transportType:type_name -> orderservice.TransportType
	3, // 4: orderservice.Order.status:type_name -> orderservice.Status
	7, // 5: orderservice.Order.GoodsList:type_name -> orderservice.OrderRow
	4, // 6: orderservice.OrderService.Get:input_type -> orderservice.GetRequest
	5, // 7: orderservice.OrderService.Get:output_type -> orderservice.GetResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRow); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		EnumInfos:         file_orders_proto_enumTypes,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}
