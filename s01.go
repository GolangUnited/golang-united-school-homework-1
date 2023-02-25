package structs

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

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) SetLastName(lastName string) {
	user.lastName = lastName
}

func (user *User) FullName() string {
	return user.lastName + " " + user.firstName
}

func ResetUser(user *User) {
	user.firstName = ""
	user.lastName = ""
}

func IsUser(input interface{}) bool {
	switch input.(type) {
	case User:
		return true
	default:
		return false
	}
}

func ProcessUser(input UserInterface) string {
	return input.FullName()
}
