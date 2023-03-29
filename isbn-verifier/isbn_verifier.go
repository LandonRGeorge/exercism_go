package isbn

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	re := regexp.MustCompile(`^\d[0-9-]*?[0-9X]$`)
	if !re.MatchString(isbn) {
		return false
	}
	strippedISBN := strings.ReplaceAll(isbn, "-", "")
	if utf8.RuneCountInString(strippedISBN) != 10 {
		return false
	}
	var total int
	for i, r := range strippedISBN {
		var value int
		if r == 'X' {
			value = 10
		} else {
			value, _ = strconv.Atoi(string(r))
		}
		multiplyBy := 10 - i
		total += value * multiplyBy
	}
	return total%11 == 0
}
