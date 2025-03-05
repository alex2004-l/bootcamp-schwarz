func reverseString(s string) string {
	result := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = s[j], s[i]
	}

	return string(result)
}

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

func mirrorWords(s string) string {
	result := make([]byte, len(s))
	words := getWords(s)
	idx := 0

	if len(words) == 0 {
		return ""
	}

	for i := 0; i < len(words); i++ {
		reversed := reverseString(words[i])
		for j := 0; j < len(reversed); j++ {
			result[idx] = reversed[j]
			idx++
		}
		if i != len(words)-1 {
			result[idx] = ' '
			idx++
		}
	}

	return string(result)
}