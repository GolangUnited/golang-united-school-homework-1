package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func Hello() string {
	Message := emoji.Sprint("Hello :world_map:!")
	return Message
}

func main() {
	fmt.Println(Hello())
}
