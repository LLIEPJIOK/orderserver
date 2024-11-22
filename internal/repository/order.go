package repository

import (
	"context"
	"errors"
	"fmt"

	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/pkg/database"
)

type DB interface {
	AddOrder(ctx context.Context, item string, quantity int32) (*models.Order, error)
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
	UpdateOrder(ctx context.Context, id string, item string, quantity int32) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) (*models.Order, error)
}

type Cache interface {
	SetOrder(ctx context.Context, order *models.Order) error
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type OrderRepository struct {
	db    DB
	cache Cache
}

func NewOrderRepository(db DB, cache Cache) *OrderRepository {
	return &OrderRepository{
		db:    db,
		cache: cache,
	}
}

func (r *OrderRepository) CreateOrder(
	ctx context.Context,
	item string,
	quantity int32,
) (*models.Order, error) {
	ord, err := r.db.AddOrder(ctx, item, quantity)
	if err != nil {
		return nil, fmt.Errorf("database.AddOrder: %w", err)
	}

	err = r.cache.SetOrder(ctx, ord)
	if err != nil {
		return nil, fmt.Errorf("cache.SetOrder: %w", err)
	}

	return ord, nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	ord, err := r.cache.GetOrder(ctx, id)
	if err == nil && !errors.As(err, &database.ErrNotExists{}) {
		return ord, nil
	} else if !errors.As(err, &database.ErrNotExists{}) {
		return nil, fmt.Errorf("cache.GetOrder: %w", err)
	}

	ord, err = r.db.GetOrder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("database.GetOrder: %w", err)
	}

	return ord, nil
}

func (r *OrderRepository) UpdateOrder(
	ctx context.Context,
	id, item string,
	quantity int32,
) (*models.Order, error) {
	ord, err := r.db.UpdateOrder(ctx, id, item, quantity)
	if err != nil {
		return nil, fmt.Errorf("database.UpdateOrder: %w", err)
	}

	err = r.cache.SetOrder(ctx, ord)
	if err != nil {
		return nil, fmt.Errorf("cache.SetOrder: %w", err)
	}

	return ord, nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id string) (*models.Order, error) {
	ord, err := r.db.DeleteOrder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("database.UpdateOrder: %w", err)
	}

	err = r.cache.DeleteOrder(ctx, id)
	if err != nil && !errors.As(err, &database.ErrNotExists{}) {
		return nil, fmt.Errorf("cache.DeleteOrder: %w", err)
	}

	return ord, nil
}

func (r *OrderRepository) ListOrders(ctx context.Context) ([]*models.Order, error) {
	ord, err := r.db.ListOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("database.ListOrder: %w", err)
	}

	return ord, nil
}
