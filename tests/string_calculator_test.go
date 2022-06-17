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
