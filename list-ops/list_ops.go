package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accum := initial
	for i := 0; i < len(s); i++ {
		nbr := s[i]
		accum = fn(accum, nbr)
	}
	return accum
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accum := initial
	for i := len(s) - 1; i >= 0; i-- {
		nbr := s[i]
		accum = fn(nbr, accum)
	}
	return accum
}

func (s IntList) Filter(fn func(int) bool) IntList {
	s2 := IntList{}
	for _, j := range s {
		if !fn(j) {
			continue
		}
		s2 = append(s2, j)
	}
	return s2
}

func (s IntList) Length() int {
	var i int
	for range s {
		i++
	}
	return i
}

func (s IntList) Map(fn func(int) int) IntList {
	for i := range s {
		s[i] = fn(s[i])
	}
	return s
}

func (s IntList) Reverse() IntList {
	sLen := len(s)
	s2 := make(IntList, sLen)
	for i := 0; i < sLen; i++ {
		idx := sLen - 1 - i
		s2[i] = s[idx]
	}
	return s2
}

func (s IntList) Append(lst IntList) IntList {
	a := append(s, lst...)
	return a
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = append(s, list...)
	}
	return s
}
