package main

import "fmt"

func main() {
	var itemCount int
	fmt.Print("Enter the number of items sold: ")
	fmt.Scan(&itemCount)

	totalSales := 0.0

	for counter := 1; counter <= itemCount; counter++ {
		var itemValue float64
		fmt.Printf("Enter the value of item %d: $", counter)
		fmt.Scan(&itemValue)
		totalSales += itemValue
	}

	earnings := 200 + (0.09 * totalSales)

	fmt.Printf("Total earnings: $%.2f\n", earnings)
}
