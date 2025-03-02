func isAnagram(s string, t string) bool {
	myMap := make(map[rune]int)

	for _, c := range s {
		myMap[c]++
	}

	for _, c := range t {
		myMap[c]--
	}

	for _, v := range myMap {
		if v != 0 {
			return false
		}
	}

	return true
}