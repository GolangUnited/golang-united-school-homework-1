package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	message := emoji.Sprint("Hello :world_map:!")
	fmt.Println(message)
}
