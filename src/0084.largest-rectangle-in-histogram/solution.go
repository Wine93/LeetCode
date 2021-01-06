package solution

/*
输入: [2,1,5,6,2,3]
输出: 10
*/

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights)

	for i, h := range heights {
		l, r := i, i

		for l >= 0 && heights[l] >= h {
			l--
		}

		for r < n && heights[r] >= h {
			r++
		}

		if (r-l-1)*h >= maxArea {
			maxArea = (r - l - 1) * h
		}
	}

	return maxArea
}
