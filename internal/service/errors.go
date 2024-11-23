package service

type ErrNegativeQuantity struct{}

func NewErrNegativeQuantity() error {
	return ErrNegativeQuantity{}
}

func (e ErrNegativeQuantity) Error() string {
	return "quantity should be non-negative"
}

type ErrTimeout struct{}

func NewErrTimeout() error {
	return ErrTimeout{}
}

func (e ErrTimeout) Error() string {
	return "operation is too long"
}
