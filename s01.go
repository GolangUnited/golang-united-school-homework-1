package structs

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}
