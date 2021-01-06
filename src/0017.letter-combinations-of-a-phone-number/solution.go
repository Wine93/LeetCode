package solution

/*
输入："23"
输出：[].
*/

var (
	num2letter map[byte][]string
	ret        []string
	end        int
	number     string
)

func dfs(idx int, s string) {
	if idx == end {
		ret = append(ret, s)
		return
	}

	letters := num2letter[number[idx]]
	for _, letter := range letters {
		dfs(idx+1, s+letter)
	}

}

func letterCombinations(digits string) []string {
	num2letter = map[byte][]string{
		'1': {"!", "@", "#"},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	ret = make([]string, 0)
	end = len(digits)
	number = digits

	dfs(0, "")

	return ret
}
