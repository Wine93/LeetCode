package solution

import (
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

var (
	assert *Assert.Assertions
)

type input struct {
	n int
}

type output struct {
	out []string
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
		{
			input{15},
			output{
				[]string{
					"1",
					"2",
					"Fizz",
					"4",
					"Buzz",
					"Fizz",
					"7",
					"8",
					"Fizz",
					"Buzz",
					"11",
					"Fizz",
					"13",
					"14",
					"FizzBuzz",
				},
			},
		},
	}

	for _, t := range ts {
		input := t.input
		output := t.output

		equal(input, fizzBuzz(input.n), output.out)
	}
}
