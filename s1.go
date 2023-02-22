package structs

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

type User struct {
	FirstName string
	LastName  string
}

func New() *User {
	return &User{}
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func ResetUser(ui UserInterface) {
	ui.SetFirstName("")
	ui.SetLastName("")
}

func IsUser(ui interface{}) bool {
	_, ok := ui.(*User)
	return ok
}

func ProcessUser(ui UserInterface) string {
	if !IsUser(ui) {
		return ""
	}
	ui.SetFirstName("Jane")
	ui.SetLastName("Doe")
	return ui.FullName()
}
