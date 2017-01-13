package main

import "fmt"

func main() {
	hello := make(chan string)

	go func() {
		hello <- "Hello"
		close(hello)
	}()

	var input string
	fmt.Scanln(&input)

	for elem := range hello {
		fmt.Println(elem)
	}
	
	fmt.Println("End.")
}
