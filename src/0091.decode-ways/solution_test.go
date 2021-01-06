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
	s string
}

type output struct {
	count int
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	initTest(t)

	ts := []testCase{
		{input{"12"}, output{2}},
	}

	for _, t := range ts {
		input := t.input
		output := t.output

		equal(input, numDecodings(input.s), output.count)
	}
}
