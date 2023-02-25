package structs

import (
	"fmt"
	"reflect"
)

type User struct {
	firstName string
	lastName  string
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) SetLastName(lastName string) {
	user.lastName = lastName
}

func (user *User) FullName() string {
	return fmt.Sprintf("%s %s", user.firstName, user.lastName)
}

func New() *User {
	return new(User)
}

func ResetUser(input UserInterface) {
	input.SetFirstName("")
	input.SetLastName("")
}

func IsUser(input UserInterface) bool {
	return reflect.TypeOf(input) == reflect.TypeOf("structs.User")
}

func ProcessUser(input UserInterface) string {
	return input.FullName()
}
