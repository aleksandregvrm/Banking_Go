package main

import "fmt"

// structs "example.com/banking/structs"

// fileOps "example.com/banking/fileops"
// "github.com/Pallinder/go-randomdata"

const accountBalanceFile = "balance.txt"

type customString string

func (text customString) name(parameter string) {
	fmt.Println(text, parameter)
}

func main() {
	var newVar customString
	newVar = "dachko"
	newVar.name("something else")

	// structs.SetStruct()
	// pointers.Pointer()
	// var accountBalance, err = fileOps.GetFloatFromFile(accountBalanceFile)

	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err)
	// 	fmt.Println("---------")
	// 	// panic("Can't continue, sorry.")
	// }

	// fmt.Println("Welcome to Go Bank!")
	// fmt.Println(randomdata.PhoneNumber())

	// for {
	// 	greetWithOptions()

	// 	var choice int
	// 	fmt.Print("Your choice: ")
	// 	fmt.Scan(&choice)

	// 	// wantsCheckBalance := choice == 1

	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Your balance is", accountBalance)
	// 	case 2:
	// 		fmt.Print("Your deposit: ")
	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)

	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			// return
	// 			continue
	// 		}

	// 		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 		fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
	// 	case 3:
	// 		fmt.Print("Withdrawal amount: ")
	// 		var withdrawalAmount float64
	// 		fmt.Scan(&withdrawalAmount)

	// 		if withdrawalAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			continue
	// 		}

	// 		if withdrawalAmount > accountBalance {
	// 			fmt.Println("Invalid amount. You can't withdraw more than you have.")
	// 			continue
	// 		}

	// 		accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 		fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
	// 	default:
	// 		fmt.Println("Goodbye!")
	// 		fmt.Println("Thanks for choosing our bank")
	// 		return
	// 		// break
	// 	}
	// }
}
