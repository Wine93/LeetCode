package solution

import (
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

var (
	assert *Assert.Assertions
)

func initTest(t *testing.T) {
	assert = Assert.New(t)
}

func equal(input, output, expected interface{}) bool {
	assert.Equal(output, expected,
		"输入: %v\n输出: %v\n预期结果: %v\n",
		input, output, expected)

	return true
}

type input struct {
	heights []int
}

type output struct {
	one int
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	initTest(t)

	ts := []testCase{
		{
			input{[]int{2, 1, 5, 6, 2, 3}},
			output{10},
		},
	}

	for _, t := range ts {
		input := t.input
		output := t.output

		equal(input, largestRectangleArea(input.heights), output.one)
	}
}
