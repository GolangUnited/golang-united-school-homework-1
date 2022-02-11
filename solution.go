package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	helloWorld := emoji.Sprint("Hello :world_map:!")
	return helloWorld
}
