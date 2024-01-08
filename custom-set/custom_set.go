package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]struct{}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, elem := range l {
		s.Add(elem)
	}
	return s
}

func (s Set) String() string {
	var arr []string
	for k := range s {
		arr = append(arr, fmt.Sprintf("%q", k))
	}
	arrStr := strings.Join(arr, ", ")
	return "{" + arrStr + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	var sLarge Set
	var sSmall Set

	if len(s1) >= len(s2) {
		sLarge, sSmall = s1, s2
	} else {
		sLarge, sSmall = s2, s1
	}

	for kLarge := range sLarge {
		_, ok := sSmall[kLarge]
		if ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	var sLarge Set
	var sSmall Set

	if len(s1) >= len(s2) {
		sLarge, sSmall = s1, s2
	} else {
		sLarge, sSmall = s2, s1
	}

	s := New()
	for kLarge := range sLarge {
		if _, ok := sSmall[kLarge]; ok {
			s.Add(kLarge)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			s.Add(k1)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	sAll := New()
	ss := []Set{s1, s2}
	for _, s := range ss {
		for e := range s {
			sAll.Add(e)
		}
	}
	return sAll
}
