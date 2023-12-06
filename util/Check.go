package util

func RuneIsNumber(input rune) bool {
	return int(input) >= int('0') && int(input) <= int('9')
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}