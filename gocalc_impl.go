package gocalc

// TODO mengimplementasikan pengecekan kesalahan jika terdapat error saat melakukan operasi matematika
// mungkin berdasarkan dari nilai parameternya atau jenis tipe data

// Operasi penjumlahan
func (c *Chainable[T]) Plus(n T) IChainable[T] {
	c.lastOperation = Plus
	if c.IsError() {
		return c
	}
	result := c.sums + n
	return c.markResult(result)
}

// Operasi pengurangan
func (c *Chainable[T]) Min(n T) IChainable[T] {
	c.lastOperation = Min
	if c.IsError() {
		return c
	}
	result := c.sums - n
	return c.markResult(result)
}

// Operasi perkalian
func (c *Chainable[T]) Multiply(n T) IChainable[T] {
	c.lastOperation = Multiply
	if c.IsError() {
		return c
	}
	result := c.sums * n
	return c.markResult(result)
}

// Operasi pembagian
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
