package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var age int = 26
	const isCool = true

	// Shorthand
	name := "Cody"
	size := 1.3

	fmt.Println(name, age)
	fmt.Printf("%T\n", size)
}
