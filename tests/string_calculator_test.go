package tests

import (
	"testing"

	sc "github.com/ekinbulut/tdd-kata-go/internals"
	"github.com/stretchr/testify/assert"
)

func Test_Add_With_Empty_String(t *testing.T) {

	c := sc.NewStringCalculator()
	res, _ := c.Add("")
	assert.Equal(t, 0, res)
}

func Test_Add_With_String(t *testing.T) {

	c := sc.NewStringCalculator()
	res, _ := c.Add("1")
	assert.Equal(t, 1, res)
	res, _ = c.Add("1,2")
	assert.Equal(t, 3, res)
	res, _ = c.Add("1\n2,3")
	assert.Equal(t, 6, res)
	res, _ = c.Add("//;\n1;2")
	assert.Equal(t, 3, res)
}

func Test_Add_Negative_Numbers_Not_Allowed(t *testing.T) {
	c := sc.NewStringCalculator()
	_, err := c.Add("-1,-2")

	assert.Error(t, err, "negatives not allowed: -1,-2")
}
