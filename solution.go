package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map:!")
	fmt.Println(message)
	return message
}
