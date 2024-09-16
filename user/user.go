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
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
			CreatedAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birth date are required!")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}
