package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	match, _ := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)].*$`, text)
	return match
}

func SplitLogLine(text string) []string {
	if len(text) == 0 {
		return []string{""}
	}
	if text[0] == '<' && text[len(text)-1] == '>' {
		return []string{text}
	}
	return regexp.MustCompile("<.*?>").Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	for _, line := range lines {
		match, _ := regexp.MatchString(`(?i)".*?password.?"`, line)
		if !match {
			continue
		}
		count++
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func replace(text string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(text)
	if match == nil {
		return text
	}
	user := match[1]
	return fmt.Sprintf("[USR] %s %s", user, text)
}

func TagWithUserName(lines []string) []string {
	newLines := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		newLine := replace(line, re)
		newLines[i] = newLine
	}
	return newLines
}
