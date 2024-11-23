package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/service"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name           string
		item           string
		quantity       int32
		repoBuilder    func(t *testing.T) service.OrderRepository
		expectedResult *models.Order
		expectedErr    error
	}{
		{
			name:     "CreateOrder success",
			item:     "TestItem",
			quantity: 5,
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("CreateOrder", mock.Anything, "TestItem", int32(5)).
					Return(&models.Order{ID: "1", Item: "TestItem", Quantity: 5}, nil).
					Once()
				return mockRepo
			},
			expectedResult: &models.Order{ID: "1", Item: "TestItem", Quantity: 5},
			expectedErr:    nil,
		},
		{
			name:     "CreateOrder negative quantity",
			item:     "TestItem",
			quantity: -1,
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				return mockRepo
			},
			expectedResult: nil,
			expectedErr:    service.NewErrNegativeQuantity(),
		},
		{
			name:     "CreateOrder error from repo",
			item:     "TestItem",
			quantity: 5,
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("CreateOrder", mock.Anything, "TestItem", int32(5)).
					Return(nil, assert.AnError).
					Once()
				return mockRepo
			},
			expectedResult: nil,
			expectedErr:    assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoBuilder(t)
			orderService := service.NewOrderService(repo)

			result, err := orderService.CreateOrder(context.Background(), tt.item, tt.quantity)

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	tests := []struct {
		name           string
		repoBuilder    func(t *testing.T) service.OrderRepository
		orderID        string
		expectedResult *models.Order
		expectedErr    error
	}{
		{
			name: "GetOrder success",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("GetOrder", mock.Anything, "1").
					Return(&models.Order{ID: "1", Item: "TestItem", Quantity: 2}, nil).
					Once()
				return mockRepo
			},
			orderID: "1",
			expectedResult: &models.Order{
				ID:       "1",
				Item:     "TestItem",
				Quantity: 2,
			},
			expectedErr: nil,
		},
		{
			name: "GetOrder not found",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("GetOrder", mock.Anything, "nonexistent").
					Return(nil, assert.AnError).
					Once()
				return mockRepo
			},
			orderID:        "nonexistent",
			expectedResult: nil,
			expectedErr:    assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoBuilder(t)
			orderService := service.NewOrderService(repo)

			result, err := orderService.GetOrder(context.Background(), tt.orderID)

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	tests := []struct {
		name           string
		repoBuilder    func(t *testing.T) service.OrderRepository
		id             string
		item           string
		quantity       int32
		expectedResult *models.Order
		expectedErr    error
	}{
		{
			name: "UpdateOrder success",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("UpdateOrder", mock.Anything, "1", "UpdatedItem", int32(5)).
					Return(&models.Order{ID: "1", Item: "UpdatedItem", Quantity: 5}, nil).
					Once()
				return mockRepo
			},
			id:             "1",
			item:           "UpdatedItem",
			quantity:       5,
			expectedResult: &models.Order{ID: "1", Item: "UpdatedItem", Quantity: 5},
			expectedErr:    nil,
		},
		{
			name: "UpdateOrder negative quantity",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				return mockRepo
			},
			id:             "1",
			item:           "UpdatedItem",
			quantity:       -1,
			expectedResult: nil,
			expectedErr:    service.NewErrNegativeQuantity(),
		},
		{
			name: "UpdateOrder error from repo",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("UpdateOrder", mock.Anything, "1", "UpdatedItem", int32(5)).
					Return(nil, assert.AnError).
					Once()
				return mockRepo
			},
			id:             "1",
			item:           "UpdatedItem",
			quantity:       5,
			expectedResult: nil,
			expectedErr:    assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoBuilder(t)
			orderService := service.NewOrderService(repo)

			result, err := orderService.UpdateOrder(context.Background(), tt.id, tt.item, tt.quantity)

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	tests := []struct {
		name           string
		repoBuilder    func(t *testing.T) service.OrderRepository
		id             string
		expectedResult *models.Order
		expectedErr    error
	}{
		{
			name: "DeleteOrder success",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("DeleteOrder", mock.Anything, "1").
					Return(&models.Order{ID: "1", Item: "TestItem", Quantity: 5}, nil).
					Once()
				return mockRepo
			},
			id:             "1",
			expectedResult: &models.Order{ID: "1", Item: "TestItem", Quantity: 5},
			expectedErr:    nil,
		},
		{
			name: "DeleteOrder error from repo",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("DeleteOrder", mock.Anything, "1").
					Return(nil, assert.AnError).
					Once()
				return mockRepo
			},
			id:             "1",
			expectedResult: nil,
			expectedErr:    assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoBuilder(t)
			orderService := service.NewOrderService(repo)

			result, err := orderService.DeleteOrder(context.Background(), tt.id)

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestListOrders(t *testing.T) {
	tests := []struct {
		name           string
		repoBuilder    func(t *testing.T) service.OrderRepository
		expectedResult []*models.Order
		expectedErr    error
	}{
		{
			name: "ListOrders success",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("ListOrders", mock.Anything).
					Return([]*models.Order{
						{ID: "1", Item: "TestItem1", Quantity: 5},
						{ID: "2", Item: "TestItem2", Quantity: 3},
					}, nil).
					Once()
				return mockRepo
			},
			expectedResult: []*models.Order{
				{ID: "1", Item: "TestItem1", Quantity: 5},
				{ID: "2", Item: "TestItem2", Quantity: 3},
			},
			expectedErr: nil,
		},
		{
			name: "ListOrders error from repo",
			repoBuilder: func(t *testing.T) service.OrderRepository {
				t.Helper()
				mockRepo := NewMockOrderRepository(t)
				mockRepo.
					On("ListOrders", mock.Anything).
					Return(nil, assert.AnError).
					Once()
				return mockRepo
			},
			expectedResult: nil,
			expectedErr:    assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoBuilder(t)
			orderService := service.NewOrderService(repo)

			result, err := orderService.ListOrders(context.Background())

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
