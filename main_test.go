package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBasicSubtotalCalculation(t *testing.T) {
	// Given I have an order with 10 items at $50 each
	// When I run `pricer 10 50.00`
	// Then I see the output `500.00`
	
	cmd := exec.Command("go", "run", "main.go", "10", "50.00")
	output, err := cmd.Output()
	
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
	
	result := strings.TrimSpace(string(output))
	expected := "500.00"
	
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}