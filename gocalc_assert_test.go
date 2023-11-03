package gocalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	result := New(20).Plus(5).Divide(0).Plus(5).Multiply(100).Divide(10).Minus(3)
	t.Logf("\nlastOperation: %v\nlastSuccessOperation: %v\nlastErrorOperation: %v\nerror: %v\n",
		result.LastOperation(),
		result.LastSuccessOperation(),
		result.LastErrorOperation(),
		result.Error())

	r, err := result.ResultAndError()
	assert_.True(result.IsError())
	assert_.EqualError(err, "Can't divide number by zero")
	assert_.Equal(r, 25)
}
