package strings

import (
	"fmt"
	"strings"
)

func ManipulateStrings() {
	fmt.Println("hola"[0])
	fmt.Println(string("hola"[0]))
	fmt.Println("Cantidad de letras: ", len("hola"))

	// using strings library

	text := "Hello World: Hello Platzi: Hello GO"

	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, ":", ";", -1))
	fmt.Println(strings.Split(text, ":"))
}