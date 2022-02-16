package hello_world

import (
	"testing"
)

func TestHey(t *testing.T) {
	testCases := []struct {
		title           string
		expectedMessage string
	}{
		{
			title:           "main",
			expectedMessage: "Hello 🗺️ !",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			message := Hey()
			if message != testCase.expectedMessage {
				t.Logf("expected message: %s, got: %s", testCase.expectedMessage, message)
				t.Fail()
			}
		})
	}
}
