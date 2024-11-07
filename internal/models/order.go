package models

type Order struct {
	ID       string `json:"id" db:"id"`
	Item     string `json:"item" db:"item"`
	Quantity int32  `json:"quantity" db:"quantity"`
}
