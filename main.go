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
	
	if state == "UT" {
		tax := subtotal * 0.0685
		total := subtotal + tax
		fmt.Printf("%.2f\n", total)
	} else {
		fmt.Printf("%.2f\n", subtotal)
	}
}