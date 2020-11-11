package raindrops

import (
	"strconv"
)

var raindrops = []struct {
	factor int
	sound string
}{
	{ 3, "Pling" },
	{ 5, "Plang" },
	{ 7, "Plong" },
}

func Convert(number int) string {
	sound := ""

	for _, raindrop := range raindrops {
		if number % raindrop.factor == 0 {
			sound += raindrop.sound
		}
	}

	if (len(sound) == 0) {
		sound = strconv.Itoa(number)
	}

	return sound
}