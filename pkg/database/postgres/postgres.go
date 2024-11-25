package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/LLIEPJIOK/orderserver/internal/models"
)

type PostgresConfig struct {
	PostgresUserName string `env:"POSTGRES_USER"     env-default:"root"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-default:"123"`
	PostgresDBName   string `env:"POSTGRES_DB"       env-default:"yandex"`
	PostgresHost     string `env:"POSTGRES_HOST"     env-default:"localhost"`
	PostgresPort     string `env:"POSTGRES_PORT"     env-default:"5432"`
}

type DB struct {
	db      *sqlx.DB
	queries *Queries
}

func NewPostgres(config PostgresConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.PostgresUserName,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresHost,
		config.PostgresPort,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if _, err := db.Conn(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{
		queries: New(db),
	}, nil
}

func (db *DB) AddOrder(ctx context.Context, item string, quantity int32) (*models.Order, error) {
	arg := AddOrderParams{
		Item:     item,
		Quantity: quantity,
	}

	res, err := db.queries.AddOrder(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to add order: %w", err)
	}

	ord := dbOrderToGlobal(res)
	return &ord, nil
}

func (db *DB) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	res, err := db.queries.GetOrder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: :%w", err)
	}

	ord := dbOrderToGlobal(res)
	return &ord, nil
}

func (db *DB) ListOrders(ctx context.Context) ([]*models.Order, error) {
	res, err := db.queries.ListOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}

	ords := make([]*models.Order, len(res))
	for i, v := range res {
		ord := dbOrderToGlobal(v)
		ords[i] = &ord
	}

	return ords, nil
}

func (db *DB) UpdateOrder(ctx context.Context, id string, item string, quantity int32) (*models.Order, error) {
	arg := UpdateOrderParams{
		ID:      id,
		Column2: item,
		Column3: quantity,
	}

	res, err := db.queries.UpdateOrder(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	ord := dbOrderToGlobal(res)
	return &ord, nil
}

func (db *DB) DeleteOrder(ctx context.Context, id string) (*models.Order, error) {
	res, err := db.queries.DeleteOrder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order: %w", err)
	}

	ord := dbOrderToGlobal(res)
	return &ord, nil
}

func (db *DB) Close() error {
	if err := db.db.Close(); err != nil {
		return fmt.Errorf("failed to close postgres: %w", err)
	}

	return nil
}

func dbOrderToGlobal(ord Order) models.Order {
	return models.Order{
		ID:       ord.ID,
		Item:     ord.Item,
		Quantity: ord.Quantity,
	}
}
