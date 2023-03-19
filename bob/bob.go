package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isUpper := strings.ToUpper(remark) == remark
	hasLetters := regexp.MustCompile(`[A-Z]`).MatchString(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion && isUpper && hasLetters:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isUpper && hasLetters:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
