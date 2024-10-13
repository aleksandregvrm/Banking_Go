package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000
	// var revenue float64 = displayRevenue()
	fmt.Println("Welcome to my bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your Balance is: %.f\n", accountBalance)
	}
	if choice == 2 {
		var moneyToDeposit float64
		fmt.Print("How much money do you want to Deposit: ")
		fmt.Scan(&moneyToDeposit)
		accountBalance = accountBalance + moneyToDeposit

		fmt.Printf("You have deposited %.f\n", moneyToDeposit)
		fmt.Printf("You have %.f on your account", accountBalance)
	}
	if choice == 3 {
		var moneyToWithdraw float64
		fmt.Print("How much money do you want to Withdraw: ")
		fmt.Scan(&moneyToWithdraw)
		accountBalance = accountBalance - moneyToWithdraw

		fmt.Printf("You have deposited %.f\n", moneyToWithdraw)
		fmt.Printf("You have %.f on your account", accountBalance)
	}
	if choice == 4 {
		fmt.Println("You have exited from the program")
		return
	}
}

// func displayRevenue() float64 {
// 	var revenueToCalculate float64
// 	fmt.Print("Revenue:")
// 	fmt.Scan(&revenueToCalculate)
// 	return revenueToCalculate
// }
