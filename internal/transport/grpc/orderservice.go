package grpc

import (
	"context"
	"fmt"

	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
	client "gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/pkg/api/order"
)

type Service interface {
	CreateOrder(ctx context.Context, item string, quantity int32) (*models.Order, error)
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	UpdateOrder(ctx context.Context, id, item string, quantity int32) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) (*models.Order, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
}

type OrderService struct {
	client.UnimplementedOrderServiceServer
	service Service
}

func NewOrderService(srv Service) *OrderService {
	return &OrderService{
		service: srv,
	}
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	req *client.CreateOrderRequest,
) (*client.CreateOrderResponse, error) {
	ord, err := s.service.CreateOrder(ctx, req.GetItem(), req.GetQuantity())
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}

	return &client.CreateOrderResponse{
		Id: ord.ID,
	}, nil
}

func (s *OrderService) GetOrder(
	ctx context.Context,
	req *client.GetOrderRequest,
) (*client.GetOrderResponse, error) {
	ord, err := s.service.GetOrder(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}

	return &client.GetOrderResponse{
		Order: &client.Order{
			Id:       ord.ID,
			Item:     ord.Item,
			Quantity: ord.Quantity,
		},
	}, nil
}

func (s *OrderService) UpdateOrder(
	ctx context.Context,
	req *client.UpdateOrderRequest,
) (*client.UpdateOrderResponse, error) {
	ord, err := s.service.UpdateOrder(ctx, req.GetId(), req.GetItem(), req.GetQuantity())
	if err != nil {
		return nil, fmt.Errorf("update order: %w", err)
	}

	return &client.UpdateOrderResponse{
		Order: &client.Order{
			Id:       ord.ID,
			Item:     ord.Item,
			Quantity: ord.Quantity,
		},
	}, nil
}

func (s *OrderService) DeleteOrder(
	ctx context.Context,
	req *client.DeleteOrderRequest,
) (*client.DeleteOrderResponse, error) {
	_, err := s.service.DeleteOrder(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("delete order: %w", err)
	}

	return &client.DeleteOrderResponse{
		Success: true,
	}, nil
}

func (s *OrderService) ListOrders(
	ctx context.Context,
	req *client.ListOrdersRequest,
) (*client.ListOrdersResponse, error) {
	ordList, err := s.service.ListOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("list orders: %w", err)
	}

	grpcOrders := make([]*client.Order, 0, len(ordList))
	for _, ord := range ordList {
		grpcOrders = append(grpcOrders, &client.Order{
			Id:       ord.ID,
			Item:     ord.Item,
			Quantity: ord.Quantity,
		})
	}

	return &client.ListOrdersResponse{
		Orders: grpcOrders,
	}, nil
}
