package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func writingBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
	var accountBalance float64 = getBalanceFromFile()
	// var revenue float64 = displayRevenue()
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your Balance is: %.f\n", accountBalance)
		case 2:
			var moneyToDeposit float64
			fmt.Print("How much money do you want to Deposit: ")
			fmt.Scan(&moneyToDeposit)
			accountBalance = accountBalance + moneyToDeposit

			fmt.Printf("You have deposited %.f\n", moneyToDeposit)
			fmt.Printf("You have %.f on your account\n", accountBalance)
			writingBalanceToFile(accountBalance)
			getBalanceFromFile()
		case 3:
			var moneyToWithdraw float64
			fmt.Print("How much money do you want to Withdraw: ")
			fmt.Scan(&moneyToWithdraw)
			if moneyToWithdraw > accountBalance {
				fmt.Println("Invalid Amount!")
				continue
			}
			accountBalance = accountBalance - moneyToWithdraw

			fmt.Printf("You have withdrawn %.f\n", moneyToWithdraw)
			fmt.Printf("You have %.f on your account\n", accountBalance)
			writingBalanceToFile(accountBalance)
		default:
			fmt.Println("You have exited from the program")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}

// func displayRevenue() float64 {
// 	var revenueToCalculate float64
// 	fmt.Print("Revenue:")
// 	fmt.Scan(&revenueToCalculate)
// 	return revenueToCalculate
// }
