package structs

import "fmt"

type User struct {
	firstName string
	lastName  string
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (u *User) SetFirstName(firstname string) {
	u.firstName = firstname
}
func (u *User) SetLastName(lastname string) {
	u.lastName = lastname
}
func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func New() User {
	return User{}
}
func ResetUser(user *User) {
	user = &User{}
}
func IsUser(user User) bool {
	return fmt.Sprintf("%T\n", user) == fmt.Sprintf("%T\n", User{})
}
func ProcessUser(ui UserInterface) string {
	return ui.FullName()
}
