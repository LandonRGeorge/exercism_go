package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot string

var robots = make(map[string]int)

func makeRandom() string {
	rand.Seed(time.Now().UnixNano())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter1 := string(letters[rand.Intn(len(letters))])
	letter2 := string(letters[rand.Intn(len(letters))])
	numbers := fmt.Sprintf("%03d", rand.Intn(1000))
	return letter1 + letter2 + numbers
}

func (r *Robot) Name() (string, error) {
	if *r != "" {
		return string(*r), nil
	}
	for {
		random := makeRandom()
		_, ok := robots[random]
		if !ok {
			*r = Robot(random)
			robots[random]++
			return string(*r), nil
		}

	}
}

func (r *Robot) Reset() {
	robots = make(map[string]int)
	*r = ""
	r.Name()
}
