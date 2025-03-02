func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			good := true
			for j := 0; j < len(needle) && good; j++ {
				if needle[j] != haystack[i+j] {
					good = false
				}
			}

			if good {
				return i
			}
		}
	}

	return -1
}