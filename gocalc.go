package gocalc

type IChainable[T int | uint | float64] interface {
	// menggabungkan interface
	IChainableOperation[T]
	Sums() T
}

type IChainableOperation[T int | uint | float64] interface {
	Plus(T) IChainable[T]
	Min(T) IChainable[T]
	Multiply(T) IChainable[T]
	Divide(T) IChainable[T]
}

type Chainable[T int | uint | float64] struct {
	value T
}

func NewCalc[T int | uint | float64](value T) IChainable[T] {
	c := new(Chainable[T])
	c.value = value
	return c
}
