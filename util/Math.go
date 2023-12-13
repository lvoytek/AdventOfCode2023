package util

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return Gcd(y, x % y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func AbsInt(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func PowInt(base int, exp int) int {
	out := 1
	for i := 0; i < exp; i++ {
		out *= base
	}

	return out
}

func Sum(values []int) int {
	sum := 0
	for _, value := range(values) {
		sum += value
	}

	return sum
}