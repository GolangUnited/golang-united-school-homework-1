package structs

import (
	"fmt"
)

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

type User struct {
	firstName string
	lastName  string
}

func New() User {
	return User{}
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) FullName() string {
	return u.lastName + " " + u.firstName
}

func ResetUser(user *User) {
	(*user).firstName = ""
	(*user).lastName = ""
}

func IsUser(input interface{}) bool {
	xType := fmt.Sprintf("%T", input)
	accpectType := fmt.Sprintf("%T", User{})
	return xType == accpectType
}

func ProcessUser(input UserInterface) string {
	return input.FullName()
}
