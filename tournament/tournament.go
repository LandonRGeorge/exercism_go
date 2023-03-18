package tournament

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

type match struct {
	name    string
	outcome string
}

type team struct {
	name                string
	wins, draws, losses int
}

func (t *team) matchesPlayed() int {
	return t.wins + t.draws + t.losses
}

func (t *team) points() int {
	return (t.wins * 3) + (t.draws * 1) + (t.losses * 0)
}

func (t *team) logOutcome(outcome string) error {
	switch outcome {
	case "win":
		t.wins++
	case "draw":
		t.draws++
	case "loss":
		t.losses++
	default:
		return fmt.Errorf("%q is not a valid outcome", outcome)
	}
	return nil
}

type teamsMap map[string]team

func (tm teamsMap) toSlice() []team {
	var s []team
	for _, v := range tm {
		s = append(s, v)
	}
	sort.Slice(s, func(i int, j int) bool {
		if s[i].points() == s[j].points() {
			return s[i].name < s[j].name
		}
		return s[i].points() > s[j].points()
	})
	return s
}

func (tm teamsMap) String() string {
	var b bytes.Buffer
	ts := tm.toSlice()
	format := "%-31s|%3v |%3v |%3v |%3v |%3v\n"
	header := fmt.Sprintf(format, "Team", "MP", "W", "D", "L", "P")
	b.WriteString(header)
	for _, t := range ts {
		line := fmt.Sprintf(format, t.name, t.matchesPlayed(), t.wins, t.draws, t.losses, t.points())
		b.WriteString(line)
	}
	return b.String()
}

func getOutcome(outcome string, i int) string {
	if i == 0 {
		return outcome
	}
	m := map[string]string{
		"win":  "loss",
		"draw": "draw",
		"loss": "win",
	}
	return m[outcome]
}

func getMatches(r io.Reader) ([]match, error) {
	scanner := bufio.NewScanner(r)
	var matches []match
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || text[0] == '#' {
			continue
		}

		line := strings.Split(text, ";")
		if len(line) != 3 {
			return nil, fmt.Errorf("%v does not have 3 elements", line)
		}

		names := line[:2]
		outcome := line[2]

		for i, name := range names {
			outcome := getOutcome(outcome, i)
			m := match{
				name:    name,
				outcome: outcome,
			}
			matches = append(matches, m)
		}
	}
	return matches, nil
}

func getTeams(rs []match) (teamsMap, error) {
	var tm = make(teamsMap)
	for _, r := range rs {
		team := tm[r.name]
		if err := team.logOutcome(r.outcome); err != nil {
			return nil, err
		}
		if team.name == "" {
			team.name = r.name
		}
		tm[r.name] = team
	}
	tm.String()
	return tm, nil
}

func Tally(r io.Reader, b io.Writer) error {
	matches, err := getMatches(r)
	if err != nil {
		return err
	}
	teams, err := getTeams(matches)
	if err != nil {
		return err
	}
	if _, err := b.Write([]byte(teams.String())); err != nil {
		return err
	}
	return nil
}
