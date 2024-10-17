package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User     User
}

func (u *User) OutputUsername() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email string, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "admin",
			LastName:  "AdminGo",
			BirthDate: "____",
			CreatedAt: time.Now(),
		},
	}
}

func New(userFirstName string, userLastName string, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("firstname, lastname and birth date are required")
	}
	return &User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreatedAt: time.Now(),
	}, nil
}
