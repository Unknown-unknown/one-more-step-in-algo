package test

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"unicode"
)

func Test_125(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("isPalindrome_Regexp: %v\n", isPalindrome_Regexp(s))

	fmt.Printf("isPalindrome_Unicode: %v\n", isPalindrome_Unicode(s))
}

func isPalindrome_Regexp(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	cleanStr := bytes.ToLower([]byte(reg.ReplaceAllString(s, "")))
	length := len(cleanStr)
	for i := 0; i < length; i++ {
		if cleanStr[i] != cleanStr[length-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome_Unicode(s string) bool {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	arr := strings.FieldsFunc(strings.ToLower(s), f)
	fmt.Println(strings.Join(arr, ""))
	return true
}
