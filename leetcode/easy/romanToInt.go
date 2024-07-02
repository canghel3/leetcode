package easy

func RomanToInt(s string) int {
	links := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	coupled := map[string][]string{
		"I": {"V", "X"},
		"X": {"L", "C"},
		"C": {"D", "M"},
	}

	var result int
	for i := 0; i < len(s); i++ {
		current := string(s[i])

		var next string
		hasNext := i+1 < len(s)
		if hasNext {
			next = string(s[i+1])
		}

		var isIn = false
		for _, r := range coupled[current] {
			if next == r {
				isIn = true
				break
			}
		}

		if isIn {
			result += links[next] - links[current]
			i++
		} else {
			result += links[current]
		}
	}

	return result
}
