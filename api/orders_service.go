package api

import (
	"context"
	"orderArchive/store/postgreSQL"
)

type OrderServer struct {
	Store *postgreSQL.Store
}

func (s *OrderServer) mustEmbedUnimplementedOrderServiceServer() {
	panic("implement me")
}

func (s *OrderServer) Get(ctx context.Context, request *GetRequest) (*GetResponse, error) {

	order, err := s.Store.Order().Get(request.Id)
	if err != nil {
		return nil, err
	}

	var paymentType PaymentType
	if order.PaymentType == 1 {
		paymentType = PaymentType_CASH
	} else if order.PaymentType == 2 {
		paymentType = PaymentType_CASHLESS
	} else if order.PaymentType == 3 {
		paymentType = PaymentType_CARD
	}

	var deliveryType DeliveryType
	if order.DeliveryType == 1 {
		deliveryType = DeliveryType_DELIVERY
	} else if order.DeliveryType == 2 {
		deliveryType = DeliveryType_PICKUP
	}

	var transportType TransportType
	if order.TransportType == 1 {
		transportType = TransportType_CAR
	} else if order.TransportType == 2 {
		transportType = TransportType_RAILWAY
	} else if order.TransportType == 3 {
		transportType = TransportType_AVIA
	}

	var status Status
	status = Status_DEFAULT
	if order.Status == 0 {
		status = Status_COLLECT
	} else if order.Status == 1 {
		status = Status_PROCESSED
	} else if order.Status == 2 {
		status = Status_COLLECT
	} else if order.Status == 3 {
		status = Status_COLLECTED
	} else if order.Status == 4 {
		status = Status_SHIPPED
	} else if order.Status == 5 {
		status = Status_CLOSED
	} else if order.Status == 6 {
		status = Status_CANCELED
	} else if order.Status == 7 {
		status = Status_SUSPENDED
	} else if order.Status == 8 {
		status = Status_NOTCOMPLETED
	} else if order.Status == 9 {
		status = Status_PAYMENTEXPECTED
	}

	response := &GetResponse{
		Order: &Order{
			Id:                order.ID,
			Number:            order.Number,
			Date:              order.Date.String(),
			ShippingDate:      order.ShippingDate.String(),
			DeliveryDate:      order.DeliveryDate.String(),
			Customer:          order.Customer,
			Sum:               float32(order.Sum),
			TradePoint:        order.TradePoint,
			PaymentType:       paymentType,
			DeliveryType:      deliveryType,
			TransportType:     transportType,
			Status:            status,
			Comment:           order.Comment,
			WarehouseShipping: order.WarehouseShipping,
		},
	}

	return response, err
}
