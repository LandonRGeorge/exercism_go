package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	l1Len := len(l1)
	l2Len := len(l2)

	var lSmLen int
	var lLg, lSm []int
	var l1IsLarger bool

	if l1Len >= l2Len {
		l1IsLarger = true
		lLg, lSm = l1, l2
		lSmLen = l2Len
	} else {
		l1IsLarger = false
		lLg, lSm = l2, l1
		lSmLen = l1Len
	}

	var lSmNbrMatching int
	var currentPosOuter int
	var startPosInner int
	var currentPosInner int
	for currentPosOuter < len(lSm) && currentPosInner < len(lLg) {
		eSm := lSm[currentPosOuter]
		eLg := lLg[currentPosInner]
		if eSm != eLg {
			if l1Len == l2Len {
				return RelationUnequal
			}
			currentPosOuter = 0
			startPosInner++
			currentPosInner = startPosInner
			lSmNbrMatching = 0
			continue

		}
		currentPosOuter++
		currentPosInner++
		lSmNbrMatching++

	}
	if lSmMatches := lSmNbrMatching == lSmLen; lSmMatches {
		switch {
		case l1Len == l2Len:
			return RelationEqual
		case !l1IsLarger:
			return RelationSublist
		case l1IsLarger:
			return RelationSuperlist
		}
	}
	return RelationUnequal
}
