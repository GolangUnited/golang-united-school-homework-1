package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	greetingMessage := emoji.Sprint("Hello :world_map:!")
	return greetingMessage
}
