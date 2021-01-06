package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	ans bool
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{input{[]int{2, 3, 1, 1, 4}}, output{true}},
		{input{[]int{3, 2, 1, 0, 4}}, output{false}},
		{input{[]int{2, 0, 0}}, output{true}},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.ans
		actual := canJump(input.nums)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
