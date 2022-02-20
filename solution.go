package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println(GetMessage())
}

func GetMessage() string {
	msg := emoji.Sprint("Hello :world_map:!")
	return msg
}
