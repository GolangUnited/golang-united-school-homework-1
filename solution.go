package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map!")
	return message[0 : len(message)-1]
}
