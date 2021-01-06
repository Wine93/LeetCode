package solution

import (
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

var (
	assert *Assert.Assertions
)

type input struct {
	s string
	k int
}

type output struct {
	l int
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
		{input{"aaabb", 3}, output{3}},
		{input{"ababbc", 2}, output{5}},
		{input{"bbaaacbd", 3}, output{3}},
	}

	for _, t := range ts {
		input := t.input
		output := t.output

		equal(input, longestSubstring(input.s, input.k), output.l)
	}
}
