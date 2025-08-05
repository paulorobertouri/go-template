package calculator

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := New()
	result := calc.Add(5, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Add(5, 3) = %f; want %f", result, expected)
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := New()
	result := calc.Subtract(10, 3)
	expected := 7.0
	if result != expected {
		t.Errorf("Subtract(10, 3) = %f; want %f", result, expected)
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := New()
	result := calc.Multiply(4, 5)
	expected := 20.0
	if result != expected {
		t.Errorf("Multiply(4, 5) = %f; want %f", result, expected)
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := New()

	// Test normal division
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
	}

	// Test division by zero
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}
}
