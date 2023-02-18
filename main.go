package main

type User struct {
	firstName string
	lastName  string
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName(string)
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}
func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}
func (u *User) FullName() string {
	return u.firstName + u.lastName
}
func New() *User {
	return new(User)
}

func (u *User) ResetUser(user User) {
	u.firstName = ""
	u.lastName = ""
}
func (u *User) IsUser(a interface{}) bool {
	switch a.(type) {
	case User:
		return true
	default:
		return false
	}
}

func (u *User) ProcessUser(user UserInterface) string {
	user.SetFirstName(u.firstName)
	user.SetLastName(u.lastName)
	return u.FullName()
}
