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
			name:     "discount over $1000 without tax",
			quantity: 25,
			price:    50.00,
			state:    "CA",
			expected: 1212.50,
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