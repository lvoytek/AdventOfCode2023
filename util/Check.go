package util

func RuneIsNumber(input rune) bool {
	return int(input) >= int('0') && int(input) <= int('9')
}