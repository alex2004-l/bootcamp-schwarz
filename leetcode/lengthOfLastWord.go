func lengthOfLastWord(s string) int {
    good := false
    length := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            if good {
                return length
            }
        } else {
            length ++
            good = true
        }
    }

    return length
}