package stack

import "fmt"

func CheckBracketPair(str string) bool {
	if len(str) == 0 {
		return true
	}

	for _, c := range str {
		s := string(c)
		fmt.Println(s)
	}
	return false
}
