package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculator_Add(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", 10, -5, 5},
		{"zero values", 0, 0, 0},
		{"decimal numbers", 2.5, 1.5, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 3, 7},
		{"negative result", 3, 10, -7},
		{"negative numbers", -5, -2, -3},
		{"zero result", 5, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4, 3, 12},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -4, 3, -12},
		{"zero multiplication", 5, 0, 0},
		{"decimal numbers", 2.5, 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := New()

	t.Run("valid division", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		require.NoError(t, err)
		assert.Equal(t, 5.0, result)
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "division by zero")
	})

	t.Run("decimal division", func(t *testing.T) {
		result, err := calc.Divide(7.5, 2.5)
		require.NoError(t, err)
		assert.Equal(t, 3.0, result)
	})
}

func TestCalculator_Calculate(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		operation Operation
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"add operation", Add, 5, 3, 8, false},
		{"subtract operation", Subtract, 10, 4, 6, false},
		{"multiply operation", Multiply, 3, 4, 12, false},
		{"divide operation", Divide, 15, 3, 5, false},
		{"divide by zero", Divide, 10, 0, 0, true},
		{"invalid operation", "invalid", 5, 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Calculate(tt.operation, tt.a, tt.b)
			
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestCalculator_Power(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		base      float64
		exponent  float64
		expected  float64
		expectErr bool
	}{
		{"power of 2", 2, 3, 8, false},
		{"power of 0", 5, 0, 1, false},
		{"power of 1", 7, 1, 7, false},
		{"negative exponent", 2, -1, 0, true},
		{"decimal base", 2.5, 2, 6.25, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Power(tt.base, tt.exponent)
			
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkCalculator_Add(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

func BenchmarkCalculator_Multiply(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Multiply(float64(i), float64(i+1))
	}
}

func BenchmarkCalculator_Divide(b *testing.B) {
	calc := New()
	for i := 1; i < b.N+1; i++ {
		calc.Divide(float64(i*10), float64(i))
	}
}
