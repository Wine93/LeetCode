package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	f int
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{input{"leetcode"}, output{0}},
		{input{"loveleetcode"}, output{2}},
	}

	for i, t := range ts {
		input := t.input
		expected := t.output.f
		actual := firstUniqChar(input.s)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
