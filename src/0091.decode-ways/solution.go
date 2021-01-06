package solution

func numDecodings(s string) int {
	n := len(s)
	count := map[int]int{0: 1}

	for i := 0; i < n; i++ {
		count[i] = count[i-1]

		num := (s[i-1]-'A')*10 + (s[i] - 'A')
		if num >= 1 && num <= 26 {
			count[i] = count[i] + count[i-2]
		}
	}

	return count[n-1]
}
