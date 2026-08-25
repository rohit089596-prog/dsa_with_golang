package main

import "fmt"

func main() {
	str := "rohit"
	var s []byte

	for i := 0; i < len(str); i++ {
		ch := str[i]
		s = append(s, ch)
	}

	var ans []byte
	for len(s) > 0 {
		ch := s[len(s)-1]
		ans = append(ans, ch)
		s = s[:len(s)-1]
	}

	fmt.Println("Anser is ", string(ans))
}
