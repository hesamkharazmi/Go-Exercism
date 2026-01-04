package triangle

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b >= c && a+c >= b && b+c >= a {
		if a == b && b == c {
			return Equ
		} else {
			if a == b || a == c || b == c {
				return Iso
			}
		}
	} else {
		return NaT
	}
	return Sca
}
