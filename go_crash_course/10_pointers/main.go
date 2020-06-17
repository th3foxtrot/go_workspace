package main

import "fmt"

func main() {
	a := 5
	// Assigns b to a pointer of a
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from adddress
	fmt.Println(*b)
	// Same as doing
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// Use pointers when you are moving large sets of data
	// refer to pointers instead
}
