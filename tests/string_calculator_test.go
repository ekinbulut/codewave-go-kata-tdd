package tests

import (
	"testing"

	sc "github.com/ekinbulut/tdd-kata-go/internals"
	"github.com/stretchr/testify/assert"
)

func Test_Add_With_Empty_String(t *testing.T) {

	c := sc.NewStringCalculator()
	assert.Equal(t, 0, c.Add(""))
}

func Test_Add_With_String(t *testing.T) {

	c := sc.NewStringCalculator()
	assert.Equal(t, 1, c.Add("1"))
	assert.Equal(t, 3, c.Add("1,2"))
	assert.Equal(t, 6, c.Add("1\n2,3"))
	assert.Equal(t, 6, c.Add("//;\n1;2"))
}
