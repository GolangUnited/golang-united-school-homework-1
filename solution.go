package solution

// package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map:!")
	// ioutil.WriteFile("message.txt", []byte(message[0:len(message)-1]), 0644)
	fmt.Println(message)
	return message
}

// func main() {

// 	ioutil.WriteFile("message.txt", []byte(GetMessage()), 0644)
// }
