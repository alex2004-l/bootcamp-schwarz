func capitalize(s string) string {
	solution := make([]byte, len(s))
s
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i-1] == ' ' {
			if s[i] >= 'a' && s[i] <= 'z' {
				solution[i] = 'A' + s[i] - 'a'
			} else {
				solution[i] = s[i]
			}
		} else {
			if s[i] >= 'A' && s[i] <= 'Z' {
				solution[i] = 'a' + s[i] - 'A'
			} else {
				solution[i] = s[i]
			}
		}
	}

	return string(solution)
}