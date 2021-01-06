package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	digits string
}

type output struct {
	ret []string
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{
			input{"23"},
			output{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.ret
		actual := letterCombinations(input.digits)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
