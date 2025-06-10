package main

import (
	"fmt"
	"os"
	"strconv"
)

// Money represents a monetary amount in fractional cents (1/10000 of a dollar)
type Money struct {
	fractionalCents int64 // 1/10000 of a dollar
}

func NewMoneyFromDollars(dollars float64) Money {
	// Convert to fractional cents, rounding to nearest
	return Money{fractionalCents: int64(dollars*10000 + 0.5)}
}

func (m Money) MultiplyByRate(numerator, denominator int64) Money {
	// Multiply by fraction, rounding to nearest fractional cent
	return Money{fractionalCents: (m.fractionalCents*numerator + denominator/2) / denominator}
}

func (m Money) Add(other Money) Money {
	return Money{fractionalCents: m.fractionalCents + other.fractionalCents}
}

func (m Money) IsGreaterThan(dollars float64) bool {
	return m.fractionalCents > int64(dollars*10000)
}

func (m Money) ToDollars() float64 {
	// Round to nearest cent using integer arithmetic
	cents := (m.fractionalCents + 50) / 100 // Add 50 for rounding
	return float64(cents) / 100
}

func calculateTotal(quantity int, price float64, state string) float64 {
	return calculateTotalDetailed(quantity, price, state, false)
}

var taxRates = map[string]float64{
	"UT": 0.0685, // Utah: 6.85%
	"TX": 0.0625, // Texas: 6.25%
	"CA": 0.0825, // California: 8.25%
	"NV": 0.08,   // Nevada: 8.00%
	"AL": 0.04,   // Alabama: 4.00%
}

func calculateTotalDetailed(quantity int, price float64, state string, verbose bool) float64 {
	subtotal := NewMoneyFromDollars(float64(quantity) * price)
	
	// Apply discount based on subtotal thresholds
	discountedSubtotal := subtotal
	discountRate := 0.0
	if subtotal.IsGreaterThan(50000) {
		discountedSubtotal = subtotal.MultiplyByRate(85, 100) // 15% discount (multiply by 85/100)
		discountRate = 0.15
	} else if subtotal.IsGreaterThan(10000) {
		discountedSubtotal = subtotal.MultiplyByRate(90, 100) // 10% discount (multiply by 90/100)
		discountRate = 0.10
	} else if subtotal.IsGreaterThan(7000) {
		discountedSubtotal = subtotal.MultiplyByRate(93, 100) // 7% discount (multiply by 93/100)
		discountRate = 0.07
	} else if subtotal.IsGreaterThan(5000) {
		discountedSubtotal = subtotal.MultiplyByRate(95, 100) // 5% discount (multiply by 95/100)
		discountRate = 0.05
	} else if subtotal.IsGreaterThan(1000) {
		discountedSubtotal = subtotal.MultiplyByRate(97, 100) // 3% discount (multiply by 97/100)
		discountRate = 0.03
	}
	
	// Apply tax if state is supported
	taxRate, hasTax := taxRates[state]
	var total Money
	if hasTax {
		// Convert tax rate to fraction for integer arithmetic
		// e.g., 6.85% = 685/10000
		taxNumerator := int64(taxRate * 10000 + 0.5)
		tax := discountedSubtotal.MultiplyByRate(taxNumerator, 10000)
		total = discountedSubtotal.Add(tax)
	} else {
		total = discountedSubtotal
	}
	
	if verbose {
		fmt.Printf("Subtotal: $%.2f\n", subtotal.ToDollars())
		if discountRate > 0 {
			discountNumerator := int64(discountRate * 100 + 0.5)
			discount := subtotal.MultiplyByRate(discountNumerator, 100)
			fmt.Printf("Discount (%.0f%%): -$%.2f\n", discountRate*100, discount.ToDollars())
			fmt.Printf("After discount: $%.2f\n", discountedSubtotal.ToDollars())
		}
		if hasTax {
			taxNumerator := int64(taxRate * 10000 + 0.5)
			tax := discountedSubtotal.MultiplyByRate(taxNumerator, 10000)
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
	if _, isSupported := taxRates[state]; !isSupported {
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