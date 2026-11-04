package main

import "fmt"

func main() {
	b := [4]int{1, 2, 3}
	fmt.Println(b[3])

	for i := range 4 {
		fmt.Println(i)
	}
}
