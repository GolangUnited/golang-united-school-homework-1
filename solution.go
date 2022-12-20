package solution

// package main

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map:")
	return message
}

// func main() {
// 	fmt.Println(GetMessage())
// }
