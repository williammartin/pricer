package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBasicSubtotalCalculation(t *testing.T) {
	// Given I have an order with 10 items at $50 each in a state without tax
	// When I run `pricer 10 50.00 XX`
	// Then I see the output `500.00`
	
	cmd := exec.Command("go", "run", "main.go", "10", "50.00", "XX")
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

func TestAddStateCodeParameter(t *testing.T) {
	// Given I have an order with 10 items at $50 each in a state without tax
	// When I run `pricer 10 50.00 CA`
	// Then I see the output `500.00`
	
	cmd := exec.Command("go", "run", "main.go", "10", "50.00", "CA")
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

func TestUtahTaxCalculation(t *testing.T) {
	// Given I have an order with 10 items at $50 each in Utah
	// When I run `pricer 10 50.00 UT`
	// Then I see the output `534.25`
	
	cmd := exec.Command("go", "run", "main.go", "10", "50.00", "UT")
	output, err := cmd.Output()
	
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
	
	result := strings.TrimSpace(string(output))
	expected := "534.25"
	
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}