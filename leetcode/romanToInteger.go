func romanToInt(s string) int {
	numerals := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && numerals[s[i]] < numerals[s[i+1]] {
			result += numerals[s[i+1]] - numerals[s[i]]
			i++
		} else {
			result += numerals[s[i]]
		}
	}

	return result
}