package bottlesong

import (
	"bytes"
	"strings"
	"text/template"
)

var nbrMap = map[int]string{
	10: "Ten",
	9:  "Nine",
	8:  "Eight",
	7:  "Seven",
	6:  "Six",
	5:  "Five",
	4:  "Four",
	3:  "Three",
	2:  "Two",
	1:  "One",
	0:  "No",
}

func getVerseLines(nbr int) ([]string, error) {
	before := nbrMap[nbr]
	after := nbrMap[nbr-1]
	tmplText := `{{.Before}} green {{.Before | pluralizeBottles}} hanging on the wall,
{{.Before}} green {{.Before | pluralizeBottles}} hanging on the wall,
And if one green bottle should accidentally fall,
There'll be {{.After | toLower}} green {{.After | pluralizeBottles}} hanging on the wall.`

	funcMap := template.FuncMap{
		"toLower": strings.ToLower,
		"pluralizeBottles": func(nbr string) string {
			if nbr == "One" {
				return "bottle"
			}
			return "bottles"
		},
	}
	ts, err := template.New("").Funcs(funcMap).Parse(tmplText)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	data := struct {
		Before string
		After  string
	}{
		Before: before,
		After:  after,
	}
	if err := ts.Execute(&b, data); err != nil {
		return nil, err
	}
	text := b.String()
	lines := strings.Split(text, "\n")
	return lines, nil
}

func Recite(startBottles, takeDown int) []string {
	var allLines []string
	var bottlesProcessed int
	for bottlesProcessed != takeDown {
		verseLines, err := getVerseLines(startBottles - bottlesProcessed)
		if err != nil {
			return allLines
		}
		bottlesProcessed++
		allLines = append(allLines, verseLines...)
		if bottlesProcessed != takeDown {
			allLines = append(allLines, "")
		}
	}
	return allLines
}
