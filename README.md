OOP, structs and interfaces
---
Working with structs!
We are to create User object
```go
type User struct {
	
}
```
- This struct **must** implement interface
```go
type UserInterface interface {
    SetFirstName(string)
    SetLastName(string)
    FullName() string
}
```
- `user.FullName()` **must** return string with format "lastName firstName"
- lastName firstName must be unavailable to read/modify directly.
- `New()` constructor **must** be present for `User` struct
- Implement functions(not methods):
  - `func ResetUser(input) ` accepts 1 param as input and resets both lastName and firstName (set to empty value)
  - `func IsUser(input) bool ` accepts 1 param as input and return `true` is input is of type `User`
  - `func ProcessUser(input) string` accepts 1 param as input of type `UserInterface`, sets last/first names and returns fullname
