-- name: AddOrder :one
INSERT INTO orders (id, item, quantity)
VALUES (gen_random_uuid(), $1, $2)
RETURNING *;

-- name: GetOrder :one
SELECT * 
FROM orders
WHERE id = $1;

-- name: UpdateOrder :one
UPDATE orders
SET item = COALESCE(NULLIF($2, ''), item), 
    quantity = COALESCE(NULLIF($3, 0), quantity)
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :one
DELETE FROM orders
WHERE id = $1
RETURNING *;

-- name: ListOrders :many
SELECT * 
FROM orders;
