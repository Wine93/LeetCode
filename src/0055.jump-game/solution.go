package solution

func canJump(nums []int) bool {
	n := len(nums)

	maxStep := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 && i != n-1 {
			if maxStep <= i {
				return false
			}
		} else if nums[i]+i > maxStep {
			maxStep = nums[i] + i
		}
	}

	return true
}
