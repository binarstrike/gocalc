package gocalc

import (
	"golang.org/x/exp/constraints"
)

type Operation int

type Number interface {
	constraints.Integer | constraints.Float
}

var operationName = [...]string{
	"Nop",
	"Plus",
	"Minus",
	"Multiply",
	"Divide",
}

const (
	Nop Operation = iota + 1
	Plus
	Minus
	Multiply
	Divide
)

func (op Operation) String() string {
	return operationName[op-1]
}

type IChainable[T Number] interface {
	IChainableOperation[T]
	Result() T
	Error() error
	IsError() bool
	ResultAndError() (T, error)
	LastSuccessOperation() Operation
	LastErrorOperation() Operation
	LastOperation() Operation
}

type IChainableOperation[T Number] interface {
	Plus(T) IChainable[T]
	Minus(T) IChainable[T]
	Multiply(T) IChainable[T]
	Divide(T) IChainable[T]
}

type Chainable[T Number] struct {
	sums                 T
	lastErrorOperation   Operation
	lastSuccessOperation Operation
	lastOperation        Operation
	lastErrorCaught      error
}

func New[T Number](startValue T) IChainable[T] {
	c := new(Chainable[T])
	c.sums = startValue
	c.lastOperation = Nop
	c.lastSuccessOperation = Nop
	c.lastErrorOperation = Nop
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
