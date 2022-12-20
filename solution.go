package solution

// package main

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := "Hello" + emoji.Sprint(":world_map:")
	// message := emoji.Sprint("Hello :map:")
	// emoji.Println(message)
	// ioutil.WriteFile("message.txt", []byte(message[0:len(message)-1]), 0644)
	// emoji.Println(message[0:len(message)-1] + "!")
	return message[0 : len(message)-1]
}

// func main() {
// 	fmt.Println(GetMessage())
// }
