package solution

// package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :map:")
	// emoji.Println(message)
	// ioutil.WriteFile("message.txt", []byte(message[0:len(message)-1]), 0644)
	// emoji.Println(message[0:len(message)-1] + "!")
	fmt.Println(len(message))
	return message
}

// func main() {
// 	fmt.Println(GetMessage())
// }
