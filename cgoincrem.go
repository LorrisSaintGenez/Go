package main

/*
int incrementor(int i) {
  return ++i;
}

*/
import "C"
import "fmt"

func main() {
	var i int = 42
	j := C.incrementor(C.int(i))
	fmt.Print(j)
}
