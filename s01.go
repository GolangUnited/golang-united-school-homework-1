package structs

import "reflect"

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

func New() *User {
	return new(User)
}

func ResetUser(input UserInterface) {
	input.SetFirstName("")
	input.SetLastName("")
}

func IsUser(input UserInterface) bool {
	a := new(User)
	return reflect.TypeOf(input) == reflect.TypeOf(a)
}

func ProcessUser(input UserInterface) string {
	return input.FullName()
}
