package solution

import (
	//"io"
	//"os"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	const worldMap = ":world_map:"
	out := emoji.Sprint("Hello ", worldMap, "!")
	//io.WriteString(os.Stdout, out)
	return out
}
