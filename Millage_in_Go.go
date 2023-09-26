package main

import "fmt"

func main() {
	totalMiles := 0
	totalGallons := 0
	tripCount := 0

	fmt.Println("Gas Mileage Calculator")
	fmt.Println("Enter trip data (miles driven and gallons used) or enter -1 to quit:")

	for {
		fmt.Print("Miles driven (or -1 to quit): ")
		var miles int
		_, err := fmt.Scan(&miles)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if miles == -1 {
			break
		}

		fmt.Print("Gallons used: ")
		var gallons int
		_, err = fmt.Scan(&gallons)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if gallons == -1 {
			break
		}

		mpg := float64(miles) / float64(gallons)
		tripCount++
		totalMiles += miles
		totalGallons += gallons

		fmt.Printf("Trip %d MPG: %.2f\n", tripCount, mpg)
		combinedMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined MPG: %.2f\n", combinedMPG)
	}
}
