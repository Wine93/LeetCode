package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
}

type output struct {
	o []int
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{input{[]int{1, 1, 1, 2, 2, 3}, 2}, output{[]int{1, 2}}},
		{input{[]int{1}, 1}, output{[]int{1}}},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.o
		actual := topKFrequent(input.nums, input.k)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
