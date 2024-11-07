package repository

import "fmt"

type ErrOrderNotFound struct {
	id string
}

func NewErrOrderNotFound(id string) ErrOrderNotFound {
	return ErrOrderNotFound{
		id: id,
	}
}

func (e ErrOrderNotFound) Error() string {
	return fmt.Sprintf("order with id=%q not found", e.id)
}
