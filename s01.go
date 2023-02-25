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
	return fmt.Sprintf("%s %s", user.lastName, user.firstName)
}

func New() *User {
	return &User{}
}

func ResetUser(input *User) {
	input.SetFirstName("")
	input.SetLastName("")
}

func IsUser(input User) bool {
	return reflect.TypeOf(input) == reflect.TypeOf("structs.User")
}

func ProcessUser(input UserInterface) string {
	return input.FullName()
}
