// Package calculator provides basic mathematical operations.
package calculator

import (
	"errors"
	"fmt"
)

// Operation represents a mathematical operation.
type Operation string

const (
	Add      Operation = "add"
	Subtract Operation = "subtract"
	Multiply Operation = "multiply"
	Divide   Operation = "divide"
)

// Calculator provides mathematical operations.
type Calculator struct{}

// New creates a new Calculator instance.
func New() *Calculator {
	return &Calculator{}
}

// Calculate performs the specified operation on two numbers.
func (c *Calculator) Calculate(op Operation, a, b float64) (float64, error) {
	switch op {
	case Add:
		return c.Add(a, b), nil
	case Subtract:
		return c.Subtract(a, b), nil
	case Multiply:
		return c.Multiply(a, b), nil
	case Divide:
		return c.Divide(a, b)
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// Add returns the sum of a and b.
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between a and b.
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of a and b.
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of a and b.
// Returns an error if b is zero.
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Power returns a raised to the power of b using repeated multiplication.
// This is a simple implementation for demonstration purposes.
func (c *Calculator) Power(base, exponent float64) (float64, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	
	if exponent == 0 {
		return 1, nil
	}
	
	result := 1.0
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	
	return result, nil
}
