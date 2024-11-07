package service

type ErrNegativeQuantity struct{}

func NewErrNegativeQuantity() error {
	return ErrNegativeQuantity{}
}

func (e ErrNegativeQuantity) Error() string {
	return "quantity should be non-negative"
}
