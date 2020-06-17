package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func addOne(num *int) {
	*num = *num + 1
}

func whereIsBadGuy(guy *badGuy) {
	x := guy.pos.x
	y := guy.pos.y
	fmt.Println("(", x, ",", y, ")")
}

// If we have large structs we can refer to the pointer of that struct.

func main() {
	x := 5
	fmt.Println(x)

	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p := position{4, 2}
	badguy := badGuy{"Jabba The Hutt", 100, p}
	whereIsBadGuy(&badguy)
}
