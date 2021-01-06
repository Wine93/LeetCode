package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListNode(nums []int) *ListNode {
	dummy := new(ListNode)
	last := dummy.Next

	for i := 0; i < len(nums); i++ {
		last = new(ListNode)
		last.Val = nums[i]
		last = last.Next
	}

	return dummy.Next
}

type input struct {
	l1 *ListNode
	l2 *ListNode
}

type output struct {
	o *ListNode
}

type testCase struct {
	input
	output
}

func TestSolution(t *testing.T) {
	_assert := assert.New(t)

	ts := []testCase{
		{
			input{makeListNode([]int{2, 4, 3}), makeListNode([]int{5, 6, 4})},
			output{makeListNode([]int{7, 0, 8})},
		},
	}

	for i, t := range ts {
		input := t.input

		expected := t.output.o
		actual := addTwoNumbers(input.l1, input.l2)

		if actual == nil {
			fmt.Println("hello world")
		}

		_assert.Equal(expected, actual, "输入 [%d]: %v\n", i, input)
	}
}
