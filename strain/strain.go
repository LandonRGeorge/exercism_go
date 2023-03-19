package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var i2 Ints
	for _, nbr := range i {
		if !filter(nbr) {
			continue
		}
		i2 = append(i2, nbr)
	}
	return i2
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var i2 Ints
	for _, nbr := range i {
		if filter(nbr) {
			continue
		}
		i2 = append(i2, nbr)
	}
	return i2
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var l2 Lists
	for _, nbrs := range l {
		if !filter(nbrs) {
			continue
		}
		l2 = append(l2, nbrs)
	}
	return l2
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var s2 Strings
	for _, nbrs := range s {
		if !filter(nbrs) {
			continue
		}
		s2 = append(s2, nbrs)
	}
	return s2
}
