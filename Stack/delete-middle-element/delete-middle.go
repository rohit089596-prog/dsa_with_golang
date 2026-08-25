package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	deleteMiddle(&arr, 0, len(arr))
	fmt.Println(arr)
}

func deleteMiddle(inputStack *[]int, count int, size int) {

	if count == size/2 {
		*inputStack = (*inputStack)[:len(*inputStack)-1] //pop
		return
	}
	s := *inputStack
	num := s[len(s)-1]                      //top
	*inputStack = s[:len(s)-1]              // pop
	deleteMiddle(inputStack, count+1, size) // recursive call
	*inputStack = append(*inputStack, num)  //again append  to array
}
