package main

import "fmt"

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	userIDs := intSeq()
	altuserIDS := intSeq()
	orderIDs := intSeq()

	fmt.Println(userIDs()) // 1
	fmt.Println(userIDs()) // 2
	fmt.Println(altuserIDS())
	fmt.Println(orderIDs()) // 1
	fmt.Println(userIDs())  // 3
}
