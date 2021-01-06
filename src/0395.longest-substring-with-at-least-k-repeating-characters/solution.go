package solution

func isValid(count map[int]map[int]int, k, m int) bool {
}

func longestSubstring(s string, k int) int {
	count := make(map[int]map[int]int)

	for i := 1; i <= len(s); i++ {
		for j := 0; j < 26; j++ {
			if int(s[i-1]-'a') == j {
				count[j][i] = count[j][i-1] + 1
			} else {
				count[j][i] = count[j][i-1]
			}
		}
	}

	l, r := 1, len(s)

	for l <= r {
		m := (l + r) / 2

		if isValid(count, k, m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r
}
