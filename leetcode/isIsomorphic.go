func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	myMap1 := make(map[byte]byte)
	myMap2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, ok := myMap1[s[i]]
		if ok {
			if myMap1[s[i]] != t[i] {
				return false
			}
		}

		_, ok = myMap2[t[i]]
		if ok {
			if myMap2[t[i]] != s[i] {
				return false
			}
		}
		myMap1[s[i]] = t[i]
		myMap2[t[i]] = s[i]
	}

	return true
}