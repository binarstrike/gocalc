package gocalc

import "testing"

var c2 = New[int64](10)
var c1 = New(10.0)

// golang juga punya builtin feature untuk melakukan test benchmark yaitu sebuah pengujian untuk menguji performa atau
// kemampuan dan kecepatan aplikasi

func BenchmarkChainableOperation(b *testing.B) {
	b.Run("benchmark1", func(b *testing.B) {
		for i := float64(1); i < float64(b.N); i++ {
			c1.Plus(float64(i)).Multiply(0.000001233).Multiply(1.0000023323).
				Plus(i).Multiply(0.00000000123 + (-i - i - (-12345678))).Multiply(0.000000023323).
				Plus(i).Multiply(0.00000000123 + (-i - i - (-12345678))).Multiply(0.00323).
				Plus(i).Multiply(0.00000000123 + (-i - i - (-12345678))).Multiply(0.00023323)
		}
		b.Logf("c1 %.7f", c1.Result())
	})
	b.Run("benchmark2", func(b *testing.B) {
		for i := int64(1); i < int64(b.N); i++ {
			c2.Plus(-i * 10).Multiply(-1 * i).Divide(-1 * i).
				Plus(9999).Multiply(1 * -i).Divide(-1 * i).
				Plus(9999).Multiply(1 * -i).Divide(-1 * -i).
				Plus(9999).Multiply(1 * -i).Divide(-1 * -i)
		}
		b.Logf("c2 %v", c2.Result())
	})
}
