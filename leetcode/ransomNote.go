func canConstruct(ransomNote string, magazine string) bool {
	myMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		myMap[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		myMap[ransomNote[i]]--
		if myMap[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}