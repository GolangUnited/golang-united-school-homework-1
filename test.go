package main

import (
	"fmt"
)

func main() {
	reverseMe := "Hello world"
	bytesMe := []rune(reverseMe)
	var note []rune

	for i := len(bytesMe) - 1; i >= 0; i-- {
		note = append(note, bytesMe[i])
	}

	fmt.Println(string(note))
}
