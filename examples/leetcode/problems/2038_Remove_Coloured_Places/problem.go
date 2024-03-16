package problems

func winnerOfGame(colors string) bool {

	if len(colors) == 0 || len(colors) == 1 || len(colors) == 2 {
		return false
	}

	hasAlicePicked := false
	hasBobPicked := false

	for len(colors) >= 3 {

		// alice turn
		for i := 0; i < len(colors)-2; i++ {
			if colors[i] == 'A' && colors[i+1] == 'A' && colors[i+2] == 'A' {
				colors = colors[:i+1] + colors[i+2:]
				hasAlicePicked = true
				break
			}
			if i == len(colors)-3 && hasBobPicked {
				return false
			}
		}

		// bob turn
		for i := 0; i < len(colors)-2; i++ {
			if colors[i] == 'B' && colors[i+1] == 'B' && colors[i+2] == 'B' {
				colors = colors[:i+1] + colors[i+2:]
				hasBobPicked = true
				break
			}
			if i == len(colors)-3 && hasAlicePicked {
				return true
			}
		}

		if !hasAlicePicked && !hasBobPicked {
			return false
		}
	}

	return true
}
