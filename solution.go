package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println(GetMap())
}

func GetMap() string {
	return "Hello " + emoji.Sprint(":world_map:") + "!"
}
