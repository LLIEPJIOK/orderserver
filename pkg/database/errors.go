package database

import "fmt"

type ErrNotExists struct {
	id string
}

func NewErrNotExists(id string) error {
	return ErrNotExists{
		id: id,
	}
}

func (e ErrNotExists) Error() string {
	return fmt.Sprintf("entry with id=%q doesn't exist", e.id)
}
