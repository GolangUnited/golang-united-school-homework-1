package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	salute := emoji.Sprint("Hello :world_map:!")
	return salute
}
