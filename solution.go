package golang_united_school_homework_1

import (
	"fmt"
	"reflect"
)

type User struct {
	name     string
	lastName string
}
type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (u *User) SetFirstName(name string) {
	u.name = name
}
func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}
func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.lastName, u.name)
}
func NewUser() User {
	return User{}
}
func ResetUser(input *User) {
	input.name = ""
	input.lastName = ""
}
func IsUser(input any) bool {
	if input == nil || reflect.TypeOf(input) != reflect.TypeOf(User{}) {
		return false
	}
	return true
}
func ProcessUser(input UserInterface) string {
	return input.FullName()
}
