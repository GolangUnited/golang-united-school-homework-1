package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloMessage)
}
