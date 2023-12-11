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