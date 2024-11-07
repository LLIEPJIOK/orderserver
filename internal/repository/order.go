package repository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
)

type OrderRepository struct {
	mu     *sync.Mutex
	orders map[string]models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		mu:     &sync.Mutex{},
		orders: make(map[string]models.Order),
	}
}

func (r *OrderRepository) CreateOrder(_ context.Context, item string, quantity int32) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.NewString()

	for {
		if _, ok := r.orders[id]; !ok {
			break
		}
		id = uuid.NewString()
	}

	ord := models.Order{
		ID:       id,
		Item:     item,
		Quantity: quantity,
	}
	r.orders[id] = ord

	return &ord, nil
}

func (r *OrderRepository) GetOrder(_ context.Context, id string) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ord, ok := r.orders[id]
	if !ok {
		return nil, NewErrOrderNotFound(id)
	}

	return &ord, nil
}

func (r *OrderRepository) UpdateOrder(_ context.Context, id, item string, quantity int32) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ord, ok := r.orders[id]
	if !ok {
		return nil, NewErrOrderNotFound(id)
	}

	if item != "" {
		ord.Item = item
	}

	if quantity != 0 {
		ord.Quantity = quantity
	}

	r.orders[id] = ord

	return &ord, nil
}

func (r *OrderRepository) DeleteOrder(_ context.Context, id string) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ord, ok := r.orders[id]
	if !ok {
		return nil, NewErrOrderNotFound(id)
	}

	delete(r.orders, id)

	return &ord, nil
}

func (r *OrderRepository) ListOrders(_ context.Context) ([]*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ordList := make([]*models.Order, 0, len(r.orders))
	for _, ord := range r.orders {
		ordList = append(ordList, &ord)
	}

	return ordList, nil
}
