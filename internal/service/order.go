package service

import (
	"context"

	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, item string, quantity int32) (*models.Order, error)
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	UpdateOrder(ctx context.Context, id, item string, quantity int32) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) (*models.Order, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, item string, quantity int32) (*models.Order, error) {
	if quantity < 0 {
		return nil, NewErrNegativeQuantity()
	}

	return s.repo.CreateOrder(ctx, item, quantity)
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return s.repo.GetOrder(ctx, id)
}

func (s *OrderService) UpdateOrder(ctx context.Context, id, item string, quantity int32) (*models.Order, error) {
	if quantity < 0 {
		return nil, NewErrNegativeQuantity()
	}

	return s.repo.UpdateOrder(ctx, id, item, quantity)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id string) (*models.Order, error) {
	return s.repo.DeleteOrder(ctx, id)
}

func (s *OrderService) ListOrders(ctx context.Context) ([]*models.Order, error) {
	return s.repo.ListOrders(ctx)
}
