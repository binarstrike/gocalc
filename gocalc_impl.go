package gocalc

// Operasi penjumlahan
func (c *Chainable[T]) Plus(n T) IChainable[T] {
	c.value += n
	return c
}

// Operasi pengurangan
func (c *Chainable[T]) Min(n T) IChainable[T] {
	c.value -= n
	return c
}

// Operasi perkalian
func (c *Chainable[T]) Multiply(n T) IChainable[T] {
	c.value *= n
	return c
}

// Operasi pembagian
func (c *Chainable[T]) Divide(n T) IChainable[T] {
	c.value /= n
	return c
}

// Mengambil hasil dari semua opearsi
func (c Chainable[T]) Sums() T {
	return c.value
}
