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

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func ResetUser(ui UserInterface) {
	ui.SetFirstName("")
	ui.SetLastName("")
}

func IsUser(ui interface{}) bool {
	_, ok := ui.(User)
	return ok
}

func ProcessUser(u UserInterface) string {
	u.SetFirstName("Jane")
	u.SetLastName("Doe")
	return u.FullName()
}
