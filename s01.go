package solution

import "fmt"

type UserInterface interface {
    SetFirstName(string)
    SetLastName(string)
    FullName() string
}

type User struct {
	firstName string
	lastName  string
}

func NewUser(lastName, firstName string) *User {
	return &User{
		lastName:  lastName,
		firstName: firstName,
	}
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.lastName, u.firstName)
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func ResetUser(u *User) {
	u.SetFirstName("")
	u.SetLastName("")
}

func isUser(i interface{}) bool {
	_, ok := i.(*User)
	return ok
}

func ProcessUser(u UserInterface) string {
	u.SetLastName("Mukhammadkodir")
	u.SetFirstName("Mamadaliev")
	return u.FullName()
}
