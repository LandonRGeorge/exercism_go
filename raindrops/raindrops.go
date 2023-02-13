package raindrops

import "fmt"

func Convert(number int) string {
	var text string
	if number%3 == 0 {
		text += "Pling"
	}
	if number%5 == 0 {
		text += "Plang"
	}
	if number%7 == 0 {
		text += "Plong"
	}
	if text == "" {
		text = fmt.Sprintf("%v", number)
	}
	return text
}
