package kata

import "unicode"

func getStringValue(str string) (result int) {
    result = 0
    for _, char := range str {
        result += getCharValue(char)
    }
    return
}

func getCharValue(char rune) int {
    if !unicode.IsLetter(char) {
        return 0
    }
    if unicode.IsLower(char) {
        return int(char - 'a' + 1)
    }
    if unicode.IsUpper(char) {
        return int(char - 'A' + 1)
    }
    return 0
}

func NameValue(my_list []string) []int {
    resultList := make([]int, len(my_list))
    for i, str := range my_list {
        resultList[i] = getStringValue(str) * (i + 1)
    }
    return resultList
}
