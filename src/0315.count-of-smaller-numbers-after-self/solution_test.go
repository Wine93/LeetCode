package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
}

type output struct {
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{
			input{},
			output{},
		},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output
		actual := ...

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
