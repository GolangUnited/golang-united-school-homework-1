package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println(getMap())
}

func getMap() string {
	return emoji.Sprint("Hello :world_map:!")
}
