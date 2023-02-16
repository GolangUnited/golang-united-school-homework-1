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

func New() *User {
	return &User{}
}

func (u *User) SetFirstName(s string) {
	u.firstName = s
}

func (u *User) SetLastName(s string) {
	u.lastName = s
}

func (u User) FullName() string {
	return u.firstName + " " + u.lastName
}

func ResetUser(user UserInterface) {
	user.SetFirstName("")
	user.SetLastName("")
}

func IsUser(input interface{}) bool {
	return input.(UserInterface) != nil
}

func ProcessUser(input UserInterface) string {
	input.SetLastName("Doe")
	return input.FullName()
}
