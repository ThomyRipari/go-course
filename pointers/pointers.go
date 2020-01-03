package pointers

import "fmt"

func ManipulatePointers() {
	var a int
	a = 100

	// direccion de memoria de a

	b := &a

	fmt.Println(a, b)

	// con * accedo a direccion de memoria que contiene b, y modifico b y a

	*b = 200

	fmt.Println(a, *b)

	pointerModify(b)

	defer fmt.Println(a, *b)
}

func pointerModify(b *int) {
	*b = 300
	fmt.Println(*b)
}