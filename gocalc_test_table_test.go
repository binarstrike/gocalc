package gocalc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
		t.Run(test.name, func(t *testing.T) {
			r, err := test.fn(test.n).ResultAndError()
			require_.Nil(err)
			require_.Equal(test.expected, r)
		})
	}
}
