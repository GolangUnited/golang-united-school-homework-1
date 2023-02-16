package solution

import (
	"testing"
)

func TestFullName(t *testing.T) {
	user := NewUser("Eshmat", "Toshmatov")
	if user.FullName() != "Eshmat Toshmatov" {
		t.Errorf("Incorrect user fullname. Got: %v, need: %v", user.FullName(), "Eshmat Toshmatov")
	}
}

func TestIUserType(t *testing.T) {
	user := NewUser("Eshmat", "Toshmatov")
	if ok := IUser(user); !ok {
		t.Errorf("Incorrect user")
	}
}

func TestResetUser(t *testing.T) {
	user := NewUser("Eshmat", "Toshmatov")
	ResetUser(user)
	if user.firstName != "" && user.lastName != "" {
		t.Errorf("User reset, but: %v", user)
	}
}

func TestProcessUser(t *testing.T) {
	user := NewUser("Eshmat", "Toshmatov")
	fullName := user.FullName()
	newFullName := ProcessUser(user)

	if newFullName == fullName {
		t.Errorf("User's LName and FName was not reset")
	}
}