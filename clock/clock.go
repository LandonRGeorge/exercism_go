package clock

import (
	"strconv"
	"strings"
	"time"
)

type Clock string

func (c Clock) Parse() (int, int, error) {
	s := strings.Split(string(c), ":")
	h, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	m, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, 0, err
	}
	return h, m, err
}

func New(h, m int) Clock {
	t := time.Date(2006, 1, 1, h, m, 0, 0, time.UTC)
	return Clock(t.Format("15:04"))
}

func (c Clock) Add(m int) Clock {
	h0, m0, err := c.Parse()
	if err != nil {
		panic(err)
	}
	return New(h0, m0+m)

}

func (c Clock) Subtract(m int) Clock {
	h0, m0, err := c.Parse()
	if err != nil {
		panic(err)
	}
	return New(h0, m0-m)
}

func (c Clock) String() string {
	return string(c)
}
