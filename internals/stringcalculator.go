package internals

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/ekinbulut/tdd-kata-go/utils"
)

type StringCalculator struct {
	expression *regexp.Regexp
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{
		expression: regexp.MustCompile("^//(.)\n(.*)$"),
	}
}

func (sc *StringCalculator) Add(s string) (result int, err error) {
	if utils.IsStringEmpty(s) {
		return 0, nil
	}

	delimeter, numbers := sc.extractDelimeterAndNumbers(s)
	return sc.calculate(numbers, delimeter)
}

func (sc *StringCalculator) extractDelimeterAndNumbers(s string) (d string, n string) {

	var delimiter, numbers string
	delimiter = "\n"
	numbers = s

	match := sc.expression.FindStringSubmatch(s)
	if len(match) > 0 {
		delimiter = match[1]
		numbers = match[2]
	}
	return delimiter, numbers
}

func (sc *StringCalculator) calculate(s string, d string) (result int, err error) {
	var sum int
	ns := strings.Replace(s, d, ",", -1)
	nums := strings.Split(ns, ",")

	negatives := []string{}

	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if utils.IsNegative(n) {
			negatives = append(negatives, nums[i])
		}
		if err != nil {
			return 0, err
		}
		if !utils.IsNumberGreaterThan(n, 1000) {
			sum += n
		}
	}

	if len(negatives) > 0 {

		return 0, errors.New("negatives not allowed: " + strings.Join(negatives, ","))
	}

	return sum, nil
}
