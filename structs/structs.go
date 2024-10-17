package structs

import (
	"fmt"

	User "example.com/banking/user"
)

func SetStruct() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Use the newUser function directly
	appUser, err := User.New(userFirstName, userLastName, userBirthdate)
	admin := User.NewAdmin("some Email", "password1234")

	admin.User.OutputUsername()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUsername()
	appUser.ClearUserName()

	// ... do something awesome with that gathered data!

	appUser.OutputUsername()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
