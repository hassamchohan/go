package problems

func romanToInt(s string) int {
	result := 0
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	i := 0
	for i < len(s) {
		if i != len(s)-1 {
			// for IV and IX
			if s[i] == 'I' && s[i+1] == 'V' {
				result += 4
				i += 2
				continue
			}
			if s[i] == 'I' && s[i+1] == 'X' {
				result += 9
				i += 2
				continue
			}
			// for XL and XC
			if s[i] == 'X' && s[i+1] == 'L' {
				result += 40
				i += 2
				continue
			}
			if s[i] == 'X' && s[i+1] == 'C' {
				result += 90
				i += 2
				continue
			}
			// for CD and CM
			if s[i] == 'C' && s[i+1] == 'D' {
				result += 400
				i += 2
				continue
			}
			if s[i] == 'C' && s[i+1] == 'M' {
				result += 900
				i += 2
				continue
			}
		}

		result += m[string(s[i])]
		i += 1
	}

	return result
}
