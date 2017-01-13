package main

import "fmt"

func counter(s string) {
	for i := 0; i < 300; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	counter("call")

	go counter("goroutines")

	for i := 0; i < 300; i++ {
		fmt.Println("Other call :", i)
	}

	var input string;
	fmt.Scanln(&input);
	fmt.Println("End.");
}
