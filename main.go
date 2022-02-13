package main

import (
    "fmt"

    "github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	emojiMessage := emoji.Sprint("Hello :world_map:!")
	return fmt.Println(emojiMessage)
}

func main() {
	GetMessage()
}