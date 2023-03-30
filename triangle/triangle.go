package triangle

type Kind int

const (
	NaT Kind = iota
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
)

func (k Kind) typeof(nums ...float64) Kind {
	m := make(map[float64]struct{})
	for _, n := range nums {
		if n <= 0 {
			return NaT
		}
		m[n] = struct{}{}
	}
	switch {
	case nums[0]+nums[1] < nums[2], nums[0]+nums[2] < nums[1], nums[1]+nums[2] < nums[0]:
		return NaT
	case len(m) == 1:
		return Equ
	case len(m) == 2:
		return Iso
	default:
		return Sca
	}

}

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k.typeof(a, b, c)
}
