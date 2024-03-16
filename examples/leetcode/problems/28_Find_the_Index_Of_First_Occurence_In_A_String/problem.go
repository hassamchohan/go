package problems

func strStr2(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	count := 0
	index := 0
	for i := 0; i < len(haystack); i++ {
		if len(needle) > len(haystack)-i {
			return -1
		}
		if haystack[i] == needle[0] {
			k := i
			for j := 0; j < len(needle); j++ {
				if needle[j] == haystack[k] {
					count++
					k++
					index = i
				}
			}
			if count == len(needle) {
				return index
			}
		}
		count = 0
	}

	return -1
}

// better solution
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

// Notes

//Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
//l := s[2:5]

//This slices up to (but excluding) s[5].
//l = s[:5]

//And this slices up from (and including) s[2].
//l = s[2:]
