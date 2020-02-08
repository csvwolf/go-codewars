package kata

import (
    "math"
    "strings"
    "unicode"
)

func ShiftLetter(letter rune, shift rune) rune {
    shift = shift % 26
    if !unicode.IsLetter(letter) {
        return letter
    }
    result := letter + shift
    if unicode.IsUpper(letter) && result > 'Z' {
        result = 'A' + (result - 'Z' - 1)
        return result
    }
    if unicode.IsLower(letter) && result > 'z' {
        result = 'a' + (result - 'z' - 1)
        return result
    }
    return result
}

func UnshiftLetter(letter rune, shift rune) rune {
    shift = shift % 26
    if !unicode.IsLetter(letter) {
        return letter
    }
    result := letter - shift
    if unicode.IsUpper(letter) && result < 'A' {
        result = 'Z' - ('A' - result) + 1
        return result
    }
    if unicode.IsLower(letter) && result < 'a' {
        result = 'z' - ('a' - result) + 1
        return result
    }
    return result
}

func MovingShift(s string, shift int) []string {
    totalLen := len(s)
    startShift := shift
    offset := 0
    var result []string
    interval := int(math.Ceil(float64(totalLen) / 5.0))
    for i := 5; i > 0; i-- {
        var slice []string
        for j := offset; j < offset + interval && j < totalLen; j++ {
            letter := s[j]
            letterResult := ShiftLetter(rune(letter), rune(startShift))
            startShift += 1
            slice = append(slice, string(letterResult))
        }
        result = append(result, strings.Join(slice, ""))
        offset += interval
    }
    return result
}

func DemovingShift(arr []string, shift int) string {
    result := ""
    startShift := shift
    for _, element := range arr {
        for _, char := range element {
            charResult := UnshiftLetter(char, rune(startShift))
            startShift += 1
            result += string(charResult)
        }
    }
    return result
}
