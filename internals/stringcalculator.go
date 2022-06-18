package internals

import (
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (sc *StringCalculator) Add(s string) (result int) {
	if len(s) == 0 {
		return 0
	}

	ns := strings.Replace(s, "\n", ",", -1)
	nums := strings.Split(ns, ",")

	var sum int

	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			return 0
		}
		sum += n
	}

	return sum
}
