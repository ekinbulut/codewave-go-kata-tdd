package internals

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (sc *StringCalculator) Add(s string) (result int, err error) {
	if IsStringEmpty(s) {
		return 0, nil
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

func ReplaceAndSplit(s string, d string) (result int, err error) {
	var sum int
	ns := strings.Replace(s, d, ",", -1)
	nums := strings.Split(ns, ",")

	negatives := []string{}

	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if IsNegative(n) {
			negatives = append(negatives, nums[i])
		}
		if err != nil {
			return 0, err
		}
		if !IsNumberGreaterThan(n, 1000) {
			sum += n
		}
	}

	if len(negatives) > 0 {

		return 0, errors.New("negatives not allowed: " + strings.Join(negatives, ","))
	}

	return sum, nil
}

func IsNegative(n int) bool {

	return n < 0
}

func IsNumberGreaterThan(n int, l int) bool {

	return n > l
}
