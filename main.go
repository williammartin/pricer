package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Money represents a monetary amount in fractional cents (1/10000 of a dollar)
type Money struct {
	fractionalCents int64 // 1/10000 of a dollar
}

func NewMoneyFromDollars(dollars float64) Money {
	return Money{fractionalCents: int64(dollars * 10000 + 0.5)}
}

func (m Money) Multiply(factor float64) Money {
	return Money{fractionalCents: int64(float64(m.fractionalCents) * factor + 0.5)}
}

func (m Money) Add(other Money) Money {
	return Money{fractionalCents: m.fractionalCents + other.fractionalCents}
}

func (m Money) IsGreaterThan(dollars float64) bool {
	return m.fractionalCents > int64(dollars * 10000)
}

func (m Money) ToDollars() float64 {
	// Round to nearest cent
	return math.Round(float64(m.fractionalCents)/100) / 100
}

func calculateTotal(quantity int, price float64, state string) float64 {
	return calculateTotalDetailed(quantity, price, state, false)
}

func calculateTotalDetailed(quantity int, price float64, state string, verbose bool) float64 {
	subtotal := NewMoneyFromDollars(float64(quantity) * price)
	
	// Apply discount based on subtotal thresholds
	discountedSubtotal := subtotal
	discountRate := 0.0
	if subtotal.IsGreaterThan(50000) {
		discountedSubtotal = subtotal.Multiply(0.85) // 15% discount
		discountRate = 0.15
	} else if subtotal.IsGreaterThan(10000) {
		discountedSubtotal = subtotal.Multiply(0.90) // 10% discount
		discountRate = 0.10
	} else if subtotal.IsGreaterThan(7000) {
		discountedSubtotal = subtotal.Multiply(0.93) // 7% discount
		discountRate = 0.07
	} else if subtotal.IsGreaterThan(5000) {
		discountedSubtotal = subtotal.Multiply(0.95) // 5% discount
		discountRate = 0.05
	} else if subtotal.IsGreaterThan(1000) {
		discountedSubtotal = subtotal.Multiply(0.97) // 3% discount
		discountRate = 0.03
	}
	
	var taxRate float64
	var total Money
	
	if state == "UT" {
		taxRate = 0.0685
		tax := discountedSubtotal.Multiply(taxRate)
		total = discountedSubtotal.Add(tax)
	} else if state == "TX" {
		taxRate = 0.0625
		tax := discountedSubtotal.Multiply(taxRate)
		total = discountedSubtotal.Add(tax)
	} else if state == "CA" {
		taxRate = 0.0825
		tax := discountedSubtotal.Multiply(taxRate)
		total = discountedSubtotal.Add(tax)
	} else if state == "NV" {
		taxRate = 0.08
		tax := discountedSubtotal.Multiply(taxRate)
		total = discountedSubtotal.Add(tax)
	} else if state == "AL" {
		taxRate = 0.04
		tax := discountedSubtotal.Multiply(taxRate)
		total = discountedSubtotal.Add(tax)
	} else {
		total = discountedSubtotal
	}
	
	if verbose {
		fmt.Printf("Subtotal: $%.2f\n", subtotal.ToDollars())
		if discountRate > 0 {
			discount := subtotal.Multiply(discountRate)
			fmt.Printf("Discount (%.0f%%): -$%.2f\n", discountRate*100, discount.ToDollars())
			fmt.Printf("After discount: $%.2f\n", discountedSubtotal.ToDollars())
		}
		if taxRate > 0 {
			tax := discountedSubtotal.Multiply(taxRate)
			fmt.Printf("Tax (%.2f%%): $%.2f\n", taxRate*100, tax.ToDollars())
		}
		fmt.Printf("Total: $%.2f\n", total.ToDollars())
		return total.ToDollars()
	}
	
	return total.ToDollars()
}

func run(args []string) int {
	verbose := false
	
	// Check for --verbose flag
	for i, arg := range args {
		if arg == "--verbose" {
			verbose = true
			args = append(args[:i], args[i+1:]...)
			break
		}
	}
	
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: pricer <quantity> <price> <state> [--verbose]\n")
		return 1
	}

	quantity, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Quantity must be a number\n")
		return 1
	}
	
	if quantity < 0 {
		fmt.Fprintf(os.Stderr, "Quantity cannot be negative\n")
		return 1
	}

	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Price must be a number\n")
		return 1
	}
	
	if price < 0 {
		fmt.Fprintf(os.Stderr, "Price cannot be negative\n")
		return 1
	}

	state := args[2]
	supportedStates := []string{"UT", "TX", "CA", "NV", "AL"}
	isSupported := false
	for _, s := range supportedStates {
		if state == s {
			isSupported = true
			break
		}
	}
	if !isSupported {
		fmt.Fprintf(os.Stderr, "State code not supported. Supported states: UT, NV, TX, AL, CA\n")
		return 1
	}
	
	total := calculateTotalDetailed(quantity, price, state, verbose)
	if !verbose {
		fmt.Printf("%.2f\n", total)
	}
	
	return 0
}

func main() {
	exitCode := run(os.Args[1:])
	os.Exit(exitCode)
}