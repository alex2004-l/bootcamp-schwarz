func getPrefix(str1 string, str2 string) string {
	result := []byte{}
	minimum := len(str1)
	if len(str1) > len(str2) {
		minimum = len(str2)
	}

	for i := 0; i < minimum; i++ {
		if str1[i] == str2[i] {
			result = append(result, str1[i])
		} else {
			break
		}
	}

	return string(result)
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getPrefix(prefix, strs[i])
	}

	return prefix
}