package sieve

type nbr struct {
	val         int
	isComposite bool
}

func Sieve(limit int) []int {
	var ns []nbr
	for i := 2; i <= limit; i++ {
		ns = append(ns, nbr{val: i})
	}

	for i := 0; i < len(ns); i++ {
		n := ns[i]
		if n.isComposite {
			continue
		}
		for j := 0; j < len(ns); j++ {
			b := ns[j]
			if n.val != b.val && b.val%n.val == 0 {
				b.isComposite = true
				ns[j] = b
			}

		}
	}

	var vs []int
	for i := 0; i < len(ns); i++ {
		n := ns[i]
		if n.isComposite {
			continue
		}
		vs = append(vs, n.val)
	}
	return vs
}
