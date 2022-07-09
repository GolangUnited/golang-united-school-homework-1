package solution

import (
	"math"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}

func CalcSquare(sideLen float64, sidesNum int32) float64{

	if sidesNum == 3 {
		return (sideLen*sideLen*math.Sqrt(float64(3))) / float64(4)
	} else if sidesNum == 4{
		return sideLen * sideLen
	} else if sidesNum == 0 {
		return math.Pi * sideLen * sideLen
	} else {
		return 0
	}
}
