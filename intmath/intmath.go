package intmath

import "math/bits"

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
func Abs[T Integer](x T) T {
	if x < 0 {
		x = -x
	}
	return x
}

// Returns square root of non-negative integer x, truncated.
func Sqrt[T Integer](x T) T {
	if x < 0 {
		panic("radicand must be non-negative")
	}
	if x == 0 { // special case to avoid division by zero in main algorithm
		return 0
	}
	// Want to find r such that r² == x.
	// let f(r) = r² − x, solve f(r) = 0 with Newton-Raphson method: r <- r - \frac{f(r)}{f'(r)}.
	// I.e., r <- (r + (x / r)) / 2.
	// Due to convexity of graph of r² − x, if start with r > sqrt(x), iterations will not
	// "flip-flop": first time it becomes <= sqrt(x) must be the answer.
	// Use 2^ceil(bitlen/2) (guaranteed > sqrt(x)) as initial guess for r.
	var r T = 1 << ((bits.Len64(uint64(x)) + 1) / 2)
	for {
		q := x / r
		if r <= q { // r*r <= x
			return r
		}
		r = (r + q) / 2
	}
}

// Returns x ** p for integer x and non-negative integer p.
// Returns same type as x; it is caller's responsibility to ensure no overflow.
func Pow[T, P Integer](x T, p P) T {
	if p < 0 {
		panic("exponent must be non-negative")
	}
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

// Returns log of x, base b, truncated, where both are positive integers.
func Log[B, T Integer](b B, x T) T {
	if b < 2 {
		panic("base must be >= 2")
	}
	if x <= 0 {
		panic("argument must be > 0")
	}
	// Operations between the two potentially different types would not compile. So, ensure b <= x,
	// then convert b to same type as x.
	if uint64(b) > uint64(x) {
		return 0
	}
	tb := T(b)
	var res T
	for x >= tb {
		x /= tb
		res++
	}
	return res
}
