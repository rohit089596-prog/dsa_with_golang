package main

import "fmt"

func main() {
	fmt.Println(isValidParenthesis("{[()]}")) // true
	fmt.Println(isValidParenthesis("{[(])}")) // false
	fmt.Println(isValidParenthesis("((("))    // false
}

func isValidParenthesis(expresson string) bool {
	var s []byte

	for i := 0; i < len(expresson); i++ {
		ch := expresson[i]
		if ch == '[' || ch == '{' || ch == '(' {
			s = append(s, ch)
		} else {
			if len(s) > 0 {
				top := s[len(s)-1]
				if (ch == ']' && top == '[') || (ch == ')' && top == '(') || (ch == '}' && top == '{') {
					s = s[:len(s)-1]
				} else {
					return false
				}

			} else {
				return false
			}
		}

	}

	return len(s) == 0
}
