package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	message := emoji.Sprint("Hello :world_map:!")
	return message
}
