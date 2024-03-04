package gocalc

import (
	"testing"
)

var (
	num1 int             = 10
	num2 IChainable[int] = New(10)
)

func BenchmarkBasicCalculationWithMathOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i == 0 {
			continue
		}
		num1 = (((num1 + i) * i) - i) / i
	}
	b.Log(num1)
}

func BenchmarkBasicCalculationWithGocalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i == 0 {
			continue
		}
		num2.Plus(i).Multiply(i).Minus(i).Divide(i)
	}
	b.Log(num2.Result())
}
