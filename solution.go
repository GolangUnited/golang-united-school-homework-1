package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map:")
	// ioutil.WriteFile("message.txt", []byte(message[0:len(message)-1]), 0644)
	return message[0 : len(message)-1]
}
