package problems

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		s := strs[0][i]
		for j := 1; j <= len(strs)-1; j++ {
			if i >= len(strs[j]) || s != strs[j][i] {
				return result
			}
		}
		result += string(s)
	}
	return result
}
