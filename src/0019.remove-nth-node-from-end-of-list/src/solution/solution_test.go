package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	last := dummy

	for i := 0; i < len(nums); i++ {
		last.Next = &ListNode{nums[i], nil}
		last = last.Next
	}

	return dummy.Next
}

type input struct {
	head *ListNode
	n    int
}

type output struct {
	out *ListNode
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{
			input{makeListNode([]int{1, 2, 3, 4, 5}), 2},
			output{makeListNode([]int{1, 2, 3, 5})},
		},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.out
		actual := removeNthFromEnd(input.head, input.n)

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
