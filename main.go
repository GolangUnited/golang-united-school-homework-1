package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	PrintEmoji()
}

func PrintEmoji() {

	var greeting string
	var emojiCodeMap map[string]string
	codeMap := ":world_map:"

	emojiCodeMap = emoji.CodeMap()
	em := emojiCodeMap[codeMap]

	greeting = "Hello " + em

	fmt.Println(emoji.Sprint(greeting))

}
