package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	returnedMsg := emoji.Sprint("Hello :world_map:!")
	
	return returnedMsg
}
