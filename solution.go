package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloMessage)
}
