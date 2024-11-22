CREATE TABLE IF NOT EXISTS orders
(
  id       UUID NOT NULL,
  item     TEXT NOT NULL,
  quantity INT  NOT NULL,
  PRIMARY KEY (id)
);
