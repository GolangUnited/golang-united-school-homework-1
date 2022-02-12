package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloWorld := emoji.Sprint(":world_map:!")
	return "hello" + helloWorld

}
