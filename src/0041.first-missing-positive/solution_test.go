package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	o int
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{input{[]int{1, 2, 0}}, output{3}},
		{input{[]int{3, 4, -1, 1}}, output{2}},
		{input{[]int{7, 8, 9, 11, 12}}, output{1}},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.o
		actual := firstMissingPositive(input.nums)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
