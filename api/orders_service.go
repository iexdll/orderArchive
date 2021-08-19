package api

import (
	"context"
	"errors"
	"orderArchive/models"
	"orderArchive/store/postgreSQL"
	"strings"
)

type OrderServer struct {
	Store *postgreSQL.Store
}

func (s *OrderServer) mustEmbedUnimplementedOrderServiceServer() {
	panic("implement me")
}

func (s *OrderServer) Get(ctx context.Context, request *GetOrder) (*Order, error) {

	order, err := s.Store.Order().Get(request.Id)
	if err != nil {
		return nil, err
	}

	if order.Customer != strings.ToUpper(request.Customer) {
		return nil, errors.New("401. Заказ принадлежит другому клиенту")
	}

	if models.EmptyRef != request.TradePoint && order.TradePoint != strings.ToUpper(request.TradePoint) {
		return nil, errors.New("401. Заказ принадлежит другой торговой точке")
	}

	return toOrderGRPC(order), err
}

func (s *OrderServer) List(ctx context.Context, request *GetList) (*Orders, error) {

	ordersID, err := s.Store.Order().FindIDByCustomer(request.Customer, request.TradePoint, request.Limit, request.Skip)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	for _, id := range ordersID {
		order, err := s.Store.Order().Get(id)
		if err != nil {
			return nil, err
		}

		orders = append(orders, toOrderGRPC(order))
	}

	return &Orders{Orders: orders}, err

}

func (s *OrderServer) FindByGoods(ctx context.Context, request *GetByGoods) (*Orders, error) {

	var orders []*Order

	if len(request.Goods) == 0 {
		return &Orders{Orders: orders}, nil
	}

	ordersID, err := s.Store.Order().FindIDByGoods(request.Customer, request.TradePoint, request.Goods)
	if err != nil {
		return nil, err
	}

	for _, id := range ordersID {
		order, err := s.Store.Order().Get(id)
		if err != nil {
			return nil, err
		}

		orders = append(orders, toOrderGRPC(order))
	}

	return &Orders{Orders: orders}, err

}

func toOrderGRPC(order *models.Order) *Order {

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
	if order.Status == 1 {
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

	var goodsList []*OrderRow
	for _, item := range order.GoodsList {
		goodsList = append(goodsList, &OrderRow{
			RowNumber:   item.RowNumber,
			RowID:       item.RowID,
			Goods:       item.Goods,
			Group:       item.Group,
			Price:       float32(item.Price),
			Count:       item.Count,
			CountCancel: item.CountCancel,
			Comment:     item.Comment,
		})
	}

	return &Order{
		Id:                order.ID,
		Number:            order.Number,
		Date:              order.Date.Format("2006-01-02 15:04:05"),
		ShippingDate:      order.ShippingDate.Format("2006-01-02 15:04:05"),
		DeliveryDate:      order.DeliveryDate.Format("2006-01-02 15:04:05"),
		Customer:          order.Customer,
		Sum:               float32(order.Sum),
		TradePoint:        order.TradePoint,
		PaymentType:       paymentType,
		DeliveryType:      deliveryType,
		TransportType:     transportType,
		Status:            status,
		Comment:           order.Comment,
		WarehouseShipping: order.WarehouseShipping,
		GoodsList:         goodsList,
	}
}
