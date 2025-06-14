package main

import (
	"math"
	"os/exec"
	"strings"
	"testing"
)

// Acceptance test - black box test of the complete command-line interface
func TestPricerCommandLineInterface(t *testing.T) {
	// Given I have an order with 25 items at $50 each in Utah (subtotal = $1,250)
	// When I run `pricer 25 50.00 UT`
	// Then I see the output `1295.56`
	
	cmd := exec.Command("go", "run", "main.go", "25", "50.00", "UT")
	output, err := cmd.Output()
	
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
	
	result := strings.TrimSpace(string(output))
	expected := "1295.56"
	
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Unit tests for business logic
func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name     string
		quantity int
		price    float64
		state    string
		expected float64
	}{
		{
			name:     "basic subtotal calculation",
			quantity: 10,
			price:    50.00,
			state:    "XX",
			expected: 500.00,
		},
		{
			name:     "Utah tax calculation",
			quantity: 10,
			price:    50.00,
			state:    "UT",
			expected: 534.25,
		},
		{
			name:     "discount over $1000 with Utah tax",
			quantity: 25,
			price:    50.00,
			state:    "UT",
			expected: 1295.56,
		},
		{
			name:     "California tax calculation with discount",
			quantity: 25,
			price:    50.00,
			state:    "CA",
			expected: 1312.53,
		},
		{
			name:     "Texas tax calculation with discount",
			quantity: 25,
			price:    50.00,
			state:    "TX",
			expected: 1288.28,
		},
		{
			name:     "Nevada tax calculation with discount",
			quantity: 25,
			price:    50.00,
			state:    "NV",
			expected: 1309.50,
		},
		{
			name:     "Alabama tax calculation with discount",
			quantity: 25,
			price:    50.00,
			state:    "AL",
			expected: 1261.00,
		},
		{
			name:     "5% discount for orders over $5,000",
			quantity: 120,
			price:    50.00,
			state:    "UT",
			expected: 6090.45,
		},
		{
			name:     "7% discount for orders over $7,000",
			quantity: 150,
			price:    50.00,
			state:    "UT",
			expected: 7452.79,
		},
		{
			name:     "10% discount for orders over $10,000",
			quantity: 250,
			price:    50.00,
			state:    "UT",
			expected: 12020.63,
		},
		{
			name:     "15% discount for orders over $50,000",
			quantity: 1100,
			price:    50.00,
			state:    "UT",
			expected: 49952.38,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateTotal(tt.quantity, tt.price, tt.state)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("Expected %.2f, got %.2f", tt.expected, result)
			}
		})
	}
}