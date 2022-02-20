package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	want := "Hello 🗺️ !"
	res := GetMessage()
	if want != res {
		t.Errorf("GetMessage = %v, want = %v", res, want)
	}
}
