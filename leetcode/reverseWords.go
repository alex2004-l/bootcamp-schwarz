import "strings"

func splitWords (s string) []string {
    result := []string{}
    current_word := []byte{}

    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if len(current_word) > 0 {
                result = append(result, string(current_word))
                current_word = current_word[:0]
            }
        } else {
            current_word = append(current_word, s[i])
        }
    }

    if len(current_word) > 0 {
        result = append(result, string(current_word))
    }

    return result
}

func reverseList(strList []string) []string {
    for i, j := 0, len(strList) - 1; i < j; i, j = i + 1, j - 1 {
        strList[i], strList[j] = strList[j], strList[i]
    }

    return strList
}

func reverseWords(s string) string {
    words := reverseList(splitWords(s))
    return strings.Join(words, " ")
}