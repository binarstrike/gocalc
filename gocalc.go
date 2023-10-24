package gocalc

type Operation string

const (
	Plus     Operation = "Plus()"
	Min      Operation = "Min()"
	Multiply Operation = "Multiply()"
	Divide   Operation = "Divide()"
)

type IChainable[T any] interface {
	IChainableOperation[T]
	Result() T
	Error() error
	IsError() bool
	ResultAndError() (T, error)
	LastSuccessOperation() Operation
	LastErrorOperation() Operation
	LastOperation() Operation
}

type IChainableOperation[T any] interface {
	Plus(T) IChainable[T]
	Min(T) IChainable[T]
	Multiply(T) IChainable[T]
	Divide(T) IChainable[T]
}

type Chainable[T int | int8 | int16 | int32 | int64 | uint | uint8 |
	uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128] struct {
	sums                 T
	lastErrorOperation   Operation
	lastSuccessOperation Operation
	lastOperation        Operation
	lastErrorCaught      error
}

func New[T int | int8 | int16 | int32 | int64 | uint | uint8 |
	uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128](startValue T) IChainable[T] {
	c := new(Chainable[T])
	c.sums = startValue
	c.lastOperation = ""
	c.lastSuccessOperation = ""
	c.lastErrorOperation = ""
	c.lastErrorCaught = nil
	return c
}

func (c Chainable[T]) Result() T {
	return c.sums
}

func (c Chainable[T]) Error() error {
	return c.lastErrorCaught
}

func (c Chainable[T]) IsError() bool {
	return c.Error() != nil
}

func (c Chainable[T]) ResultAndError() (T, error) {
	return c.sums, c.lastErrorCaught
}

func (c Chainable[T]) LastSuccessOperation() Operation {
	return c.lastSuccessOperation
}

func (c Chainable[T]) LastOperation() Operation {
	return c.lastOperation
}

func (c Chainable[T]) LastErrorOperation() Operation {
	return c.lastErrorOperation
}

func (c *Chainable[T]) markError(value T, err error) *Chainable[T] {
	c.sums = value
	c.lastErrorOperation = c.lastOperation
	c.lastErrorCaught = err
	return c
}

func (c *Chainable[T]) markResult(value T) *Chainable[T] {
	c.sums = value
	c.lastSuccessOperation = c.lastOperation
	return c
}
