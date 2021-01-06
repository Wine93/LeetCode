package solution

import (
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

var (
	assert *Assert.Assertions
)

type input struct {
	nums   []int
	target int
}

type output struct {
	one []int
}

type testCase struct {
	input
	output
}

func initTest(t *testing.T) {
	assert = Assert.New(t)
}

func equal(input, output, expected interface{}) bool {
	assert.Equal(output, expected,
		"输入: %v\n输出: %v\n预期结果: %v\n",
		input, output, expected)

	return true
}

func TestSolution(t *testing.T) {
	initTest(t)

	ts := []testCase{
		{input{[]int{2, 7, 11, 15}, 9}, output{[]int{0, 1}}}
	}

	for _, t := range ts {
		input := t.input
		output := t.output

		equal(input, twoSum(input.nums, input.target), output.one)
	}
}
