package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func Message() string {
	Message := emoji.Sprint("Hello :world_map:!")
	return Message
}
func main() {
	fmt.Println(Message())
}
