package solution

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum := make(map[int]int)
	l := len(A)
	count := 0

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sum[A[i]+B[j]]++
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			count = count + sum[-(C[i]+D[j])]
		}
	}

	return count
}
