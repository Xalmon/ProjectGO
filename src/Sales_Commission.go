package main

import "fmt"

func main() {
	for {
		var accountNumber, beginningBalance, totalCharges, totalCredits, creditLimit int

		fmt.Print("Enter account number (or -1 to quit): ")
		fmt.Scan(&accountNumber)

		if accountNumber == -1 {
			break
		}

		fmt.Print("Enter balance at the beginning of the month: ")
		fmt.Scan(&beginningBalance)

		fmt.Print("Enter total charges this month: ")
		fmt.Scan(&totalCharges)

		fmt.Print("Enter total credits this month: ")
		fmt.Scan(&totalCredits)

		fmt.Print("Enter allowed credit limit: ")
		fmt.Scan(&creditLimit)

		newBalance := beginningBalance + totalCharges - totalCredits

		fmt.Printf("New balance: %d\n", newBalance)

		if newBalance > creditLimit {
			fmt.Println("Credit limit exceeded")
		}

		fmt.Println()
	}
}
