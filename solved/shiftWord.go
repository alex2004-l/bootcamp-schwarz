// Shift each word with a position to the left

func getWords(s string) []string {
	result := make([]string, 0, len(s)/2+1)
	word := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				result = result[:len(result)+1]
				result[len(result)-1] = string(word)
				word = word[:0]
			}
		} else {
			word = word[:len(word)+1]
			word[len(word)-1] = s[i]
		}
	}

	if len(word) > 0 {
		result = result[:len(result)+1]
		result[len(result)-1] = string(word)
		word = word[:0]
	}

	return result
}

func shiftWord(s string) string {
	result := make([]byte, len(s))
	words := getWords(s)
	idx := 0

	if len(words) == 0 {
		return ""
	} else if len(words) == 1 {
		return words[0]
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			result[idx] = words[i][j]
			idx++
		}
		result[idx] = ' '
		idx++
	}

	for i := 0; i < len(words[0]); i++ {
		result[idx] = words[0][i]
		idx++
	}

	return string(result)
}