package gocalc

import "testing"

func TestBasicCalculation(t *testing.T) {
	// Penambahan
	if c := NewCalc(10).Plus(10).Sums(); c != 20 {
		t.Errorf("Expected 20, but got %v", c)
	}

	// Pengurangan
	if c := NewCalc(100).Min(50).Sums(); c != 50 {
		t.Errorf("Expected 50, but got %v", c)
	}

	// Perkalian
	if c := NewCalc(10).Multiply(10).Sums(); c != 100 {
		t.Errorf("Expected 100, but got %v", c)
	}

	// Pembagian
	if c := NewCalc(50).Divide(2).Sums(); c != 25 {
		t.Errorf("Expected 25, but got %v", c)
	}
}

func TestChainableCalculation(t *testing.T) {
	if c := NewCalc(2).
		Multiply(10).
		Multiply(10).
		Divide(4).
		Plus(1).
		Sums(); c != 51 {
		t.Errorf("Expected 51, but got %v", c)
	}

	if c := NewCalc(10).
		Multiply(100).
		Divide(500).
		Multiply(2).
		Plus(2).
		Min(1).
		Sums(); c != 5 {
		t.Errorf("Expected 5, but got %v", c)
	}
}

func TestFloatNumberCalculation(t *testing.T) {
	if c := NewCalc(3.2).
		Divide(2.0).
		Multiply(10.0).
		Divide(2.0).
		Sums(); c != 8.0 {
		t.Errorf("Expected 8,0, but got %v", c)
	}

	if c := NewCalc(100.123).
		Multiply(2.0).
		Multiply(5.0).
		Divide(2.0).
		Sums(); c != 500.615 {
		t.Errorf("Expected 500,615, but got %v", c)
	}

	if c := NewCalc(10.0).
		Plus(15.0).
		Plus(5.0).
		Multiply(2.0).
		Min(20.0).
		Sums(); c != 40.0 {
		t.Errorf("Expected 20.0, but got %v", c)
	}
}
