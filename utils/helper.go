package utils

func IsNegative(n int) bool {

	return n < 0
}

func IsNumberGreaterThan(n int, l int) bool {

	return n > l
}

func IsStringEmpty(s string) bool {
	return len(s) == 0
}
