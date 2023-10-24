package gocalc

import "testing"

func TestBasicCalculation(t *testing.T) {
	// Penambahan
	if c := New(10).Plus(10).Result(); c != 20 {
		t.Errorf("Expected 20, but got %v", c)
	}

	// Pengurangan
	if c := New(100).Min(50).Result(); c != 50 {
		t.Errorf("Expected 50, but got %v", c)
	}

	// Perkalian
	if c := New(10).Multiply(10).Result(); c != 100 {
		t.Errorf("Expected 100, but got %v", c)
	}

	// Pembagian
	if c := New(50).Divide(2).Result(); c != 25 {
		t.Errorf("Expected 25, but got %v", c)
	}
}

func TestChainableCalculation(t *testing.T) {
	if c := New(2).
		Multiply(10).
		Multiply(10).
		Divide(4).
		Plus(1).
		Result(); c != 51 {
		t.Errorf("Expected 51, but got %v", c)
	}

	if c := New(10).
		Multiply(100).
		Divide(500).
		Multiply(2).
		Plus(2).
		Min(1).
		Result(); c != 5 {
		t.Errorf("Expected 5, but got %v", c)
	}
}

func TestFloatNumberCalculation(t *testing.T) {
	if c := New(3.2).
		Divide(2.0).
		Multiply(10.0).
		Divide(2.0).
		Result(); c != 8.0 {
		t.Errorf("Expected 8,0, but got %v", c)
	}

	if c := New(100.123).
		Multiply(2.0).
		Multiply(5.0).
		Divide(2.0).
		Result(); c != 500.615 {
		t.Errorf("Expected 500,615, but got %v", c)
	}

	if c := New(10.0).
		Plus(15.0).
		Plus(5.0).
		Multiply(2.0).
		Min(20.0).
		Result(); c != 40.0 {
		t.Errorf("Expected 20.0, but got %v", c)
	}
}
