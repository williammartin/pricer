package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	quantity, _ := strconv.Atoi(os.Args[1])
	price, _ := strconv.ParseFloat(os.Args[2], 64)
	state := os.Args[3]
	
	subtotal := float64(quantity) * price
	
	// Apply discount if subtotal > $1,000
	discountedSubtotal := subtotal
	if subtotal > 1000 {
		discountedSubtotal = subtotal * 0.97 // 3% discount
	}
	
	if state == "UT" {
		tax := discountedSubtotal * 0.0685
		total := discountedSubtotal + tax
		fmt.Printf("%.2f\n", total)
	} else {
		fmt.Printf("%.2f\n", discountedSubtotal)
	}
}