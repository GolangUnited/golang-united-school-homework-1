package solution

import "github.com/kyokomi/emoji"

func GetMessage() string {
	return "Hello " + emoji.Sprintf(":world_map:") + "!"
}
