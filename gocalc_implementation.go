package gocalc

// Plus performs addition operation.
func (c *Chainable[T]) Plus(n T) IChainable[T] {
	c.lastOperation = Plus
	if c.IsError() {
		return c
	}
	result := c.sums + n
	return c.markResult(result)
}

// Minus performs substraction operation.
func (c *Chainable[T]) Minus(n T) IChainable[T] {
	c.lastOperation = Minus
	if c.IsError() {
		return c
	}
	result := c.sums - n
	return c.markResult(result)
}

// Multiply performs multiplication operation.
func (c *Chainable[T]) Multiply(n T) IChainable[T] {
	c.lastOperation = Multiply
	if c.IsError() {
		return c
	}
	result := c.sums * n
	return c.markResult(result)
}

// Divide performs division operation.
func (c *Chainable[T]) Divide(n T) IChainable[T] {
	c.lastOperation = Divide
	if c.IsError() {
		return c
	}
	err := (error)(nil)
	result := func(err *error) T {
		defer catch(err)

		if n == 0 {
			panic("Can't divide number by zero")
		}
		return c.sums / n
	}(&err)

	if err != nil {
		return c.markError(c.sums, err)
	}
	return c.markResult(result)
}
