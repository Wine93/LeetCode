package solution

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := index[target-nums[i]]; ok {
			return []int{j, i}
		}
		index[nums[i]] = i
	}

	return nil
}
