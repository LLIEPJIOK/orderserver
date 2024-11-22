// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package postgres

import (
	"context"
)

const addOrder = `-- name: AddOrder :one
INSERT INTO orders (id, item, quantity)
VALUES (gen_random_uuid(), $1, $2)
RETURNING id, item, quantity
`

type AddOrderParams struct {
	Item     string
	Quantity int32
}

func (q *Queries) AddOrder(ctx context.Context, arg AddOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, addOrder, arg.Item, arg.Quantity)
	var i Order
	err := row.Scan(&i.ID, &i.Item, &i.Quantity)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :one
DELETE FROM orders
WHERE id = $1
RETURNING id, item, quantity
`

func (q *Queries) DeleteOrder(ctx context.Context, id string) (Order, error) {
	row := q.db.QueryRowContext(ctx, deleteOrder, id)
	var i Order
	err := row.Scan(&i.ID, &i.Item, &i.Quantity)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT id, item, quantity 
FROM orders
WHERE id = $1
`

func (q *Queries) GetOrder(ctx context.Context, id string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(&i.ID, &i.Item, &i.Quantity)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, item, quantity 
FROM orders
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(&i.ID, &i.Item, &i.Quantity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
SET item = COALESCE(NULLIF($2, ''), item), 
    quantity = COALESCE(NULLIF($3, 0), quantity)
WHERE id = $1
RETURNING id, item, quantity
`

type UpdateOrderParams struct {
	ID      string
	Column2 interface{}
	Column3 interface{}
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder, arg.ID, arg.Column2, arg.Column3)
	var i Order
	err := row.Scan(&i.ID, &i.Item, &i.Quantity)
	return i, err
}
