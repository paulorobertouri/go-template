// Package calculator provides simple mathematical operations.
package calculator

import "errors"

// Calculator provides basic math operations
type Calculator struct{}

// New creates a new Calculator
func New() *Calculator {
	return &Calculator{}
}

// Add returns the sum of a and b
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between a and b
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of a and b
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of a and b
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
