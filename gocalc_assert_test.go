package gocalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// membuat unit testing dengan memanfaatkan library testify/assert
func TestBasicCalculationAssert(t *testing.T) {
	assert_ := assert.New(t)

	result := New(10).Plus(10).Result()
	assert_.Equal(20, result)

	result = New(10).Minus(5).Result()
	assert_.Equal(5, result)

	result = New(100).Multiply(10).Result()
	assert_.Equal(1000, result)

	result = New(50).Divide(5).Result()
	assert_.Equal(10, result)
}

func TestDivideNumberByZero(t *testing.T) {
	assert_ := assert.New(t)

	c := New(20).Plus(5).Multiply(2).Divide(0).Plus(5).Multiply(100).Divide(10).Minus(3)

	result, err := c.ResultAndError()
	assert_.True(c.IsError())
	assert_.EqualError(err, "Can't divide number by zero")
	assert_.Equal("Divide", c.LastErrorOperation().String())
	assert_.Equal("Multiply", c.LastSuccessOperation().String())
	assert_.Equal("Minus", c.LastOperation().String())
	assert_.Equal(result, 50)

	t.Logf(`lastOperation: %v
lastSuccessOperation: %v
lastErrorOperation: %v
error: %v`,
		c.LastOperation(),
		c.LastSuccessOperation(),
		c.LastErrorOperation(),
		c.Error())
}

// Untuk melakukan test dengan test table fungsi test disini akan menggunakan sub test, dimana ini adalah sebuah
// fungsi test yang terdapat di dalam fungsi test. Dengan menggunakan method/fungsi t.Run() ini memungkinkan untuk
// menjalankan test di dalam fungsi test utama.
func TestSimpleCalculationWithTestTable(t *testing.T) {
	require_ := require.New(t)
	test_table1 := []struct {
		name     string
		fn       func(int) IChainable[int]
		n        int
		expected int
	}{
		{
			name:     "Plus",
			fn:       New(10).Plus,
			n:        5,
			expected: 15,
		},
		{
			name:     "Minus",
			fn:       New(10).Minus,
			n:        5,
			expected: 5,
		},
		{
			name:     "Multiply",
			fn:       New(10).Multiply,
			n:        5,
			expected: 50,
		},
		{
			name:     "Divide",
			fn:       New(10).Divide,
			n:        5,
			expected: 2,
		},
	}

	for _, test := range test_table1 {
		// sub test
		// parameter pertama adalah nama sub test dan parameter kedua adalah fungsi callback dengan parameter *testing.T
		t.Run(test.name, func(_ *testing.T) {
			result, err := test.fn(test.n).ResultAndError()
			require_.Nil(err)
			require_.Equal(test.expected, result)
		})
	}
}
