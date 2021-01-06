package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s []byte
}

type output struct {
	s []byte
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{input{[]byte{'h', 'e', 'l', 'l', 'o'}}, output{[]byte{'o', 'l', 'l', 'e', 'h'}}},
	}

	for i, t := range ts {
		input := t.input
		expected := t.output.s
		reverseString(input.s)
		actual := input.s

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
