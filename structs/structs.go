package structs

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) outputUsername() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func SetStruct() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User

	appUser = User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}
	appUser.outputUsername()

	// ... do something awesome with that gathered data!

	appUser.outputUsername()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
