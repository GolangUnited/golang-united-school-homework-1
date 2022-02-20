package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	msg := emoji.Sprint("Hello :world_map:!")
	fmt.Println(msg)
}
