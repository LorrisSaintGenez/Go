package main

import "fmt"

func looper(i int) int {
	for i < 10 {
		i++
		fmt.Println(i)
	}
	return i
}

func main() {
	j := looper(0)
	fmt.Print(j)
}
