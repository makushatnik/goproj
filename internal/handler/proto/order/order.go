package order

import (
	"goproj/internal/entity"
	"goproj/pkg/orderservice"
)

type orderService struct {
	orderservice.UnimplementedOrderServiceServer
	orders []entity.Order
}

func New() orderservice.OrderServiceServer {
	return &orderService{
		orders: []entity.Order{
			{OrderID: "1111", BasketID: "1111"},
			{OrderID: "2222", BasketID: "2222"},
			{OrderID: "3333", BasketID: "3333"},
			{OrderID: "4444", BasketID: "4444"},
			{OrderID: "5555", BasketID: "5555"},
		},
	}
}

func (os *orderService) GetOrders(_ *orderservice.Empty, stream orderservice.OrderService_GetOrdersServer) error {
	for _, order := range os.orders {
		res := &orderservice.Order{OrderID: order.OrderID, BasketID: order.BasketID}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
