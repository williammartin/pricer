package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	quantity, _ := strconv.Atoi(os.Args[1])
	price, _ := strconv.ParseFloat(os.Args[2], 64)
	
	subtotal := float64(quantity) * price
	fmt.Printf("%.2f\n", subtotal)
}