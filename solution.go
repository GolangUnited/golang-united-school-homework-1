package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}

func main() {
	fmt.Println(GetMessage())
}
