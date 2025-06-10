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
	subtotal := NewMoneyFromDollars(float64(quantity) * price)
	
	// Apply discount if subtotal > $1,000
	discountedSubtotal := subtotal
	if subtotal.IsGreaterThan(1000) {
		discountedSubtotal = subtotal.Multiply(0.97) // 3% discount
	}
	
	if state == "UT" {
		tax := discountedSubtotal.Multiply(0.0685)
		total := discountedSubtotal.Add(tax)
		return total.ToDollars()
	} else if state == "TX" {
		tax := discountedSubtotal.Multiply(0.0625)
		total := discountedSubtotal.Add(tax)
		return total.ToDollars()
	} else if state == "CA" {
		tax := discountedSubtotal.Multiply(0.0825)
		total := discountedSubtotal.Add(tax)
		return total.ToDollars()
	}
	
	return discountedSubtotal.ToDollars()
}

func main() {
	quantity, _ := strconv.Atoi(os.Args[1])
	price, _ := strconv.ParseFloat(os.Args[2], 64)
	state := os.Args[3]
	
	total := calculateTotal(quantity, price, state)
	fmt.Printf("%.2f\n", total)
}