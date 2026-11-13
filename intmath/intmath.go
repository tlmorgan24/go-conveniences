package intmath

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Returns absolute value of x. Note, if x is the minimum value of its type, the result overflows
// and stays negative. E.g. Abs(int8(-128)) -> -128.
func Abs[T Signed](x T) T {
	if x < 0 {
		x = -x
	}
	return x
}

// Returns square root of non-negative integer x, truncated.
func Sqrt[T Unsigned](x T) T {
	if x < 2 {
		return x // sqrt(0)=0, sqrt(1)=1
	}

	// Half-open binary search for last i such that i * i <= x.
	// I.e., first i such that i * i > x, then - 1.
	var l T = 1       // square(1) <= x for all x >= 2, making this an appropriate lower bound
	var r T = x/2 + 1 // square(x/2 + 1) > x for all x >= 2, making this an appropriate upper bound
	for l < r {
		m := (l + r) / 2
		if m <= x/m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}

// Returns x ** p for integer x and non-negative integer p.
// Returns same type as x; it is caller's responsibility to ensure no overflow.
func Pow[T Integer, P Unsigned](x T, p P) T {
	var res T = 1
	for p > 0 {
		// E.g.:
		// x ** 6 == (x ** 2) ** 3
		// x ** 7 == (x ** 6) * x
		if p%2 == 0 {
			x *= x
			p /= 2
		} else {
			res *= x
			p--
		}
	}
	return res
}

// Returns log of x, base b, truncated, where both are non-negative integers.
func Log[T Unsigned](b, x T) T {
	if b < 2 || x == 0 {
		panic("invalid argument")
	}
	var res T
	for x >= b {
		x /= b
		res++
	}
	return res
}
