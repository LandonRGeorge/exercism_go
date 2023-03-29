package resistorcolorduo

import "strconv"

var colorsMap = map[string]string{
	"black":  "0",
	"brown":  "1",
	"red":    "2",
	"orange": "3",
	"yellow": "4",
	"green":  "5",
	"blue":   "6",
	"violet": "7",
	"grey":   "8",
	"white":  "9",
}

func Value(colors []string) int {
	var concatNbrs string
	for _, color := range colors[:2] {
		concatNbrs += colorsMap[color]
	}
	value, err := strconv.Atoi(concatNbrs)
	if err != nil {
		panic(err)
	}
	return value
}
