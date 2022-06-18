package internals

import (
	"regexp"
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (sc *StringCalculator) Add(s string) (result int) {
	if IsStringEmpty(s) {
		return 0
	}

	// create a regexp to match the delimiter and the numbers
	re := regexp.MustCompile("^//(.)\n(.*)$")
	match := re.FindStringSubmatch(s)
	if len(match) > 0 {
		delimiter := match[1]
		numbers := match[2]
		return ReplaceAndSplit(numbers, delimiter)
	} else {
		return ReplaceAndSplit(s, "\n")
	}
}

func IsStringEmpty(s string) bool {
	return len(s) == 0
}

func ReplaceAndSplit(s string, d string) (result int) {
	var sum int
	ns := strings.Replace(s, d, ",", -1)
	nums := strings.Split(ns, ",")

	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			return 0
		}
		sum += n
	}

	return sum
}
