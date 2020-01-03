package routines

import "fmt"

func helloGo(index int) {
	fmt.Println("Hola #", index)
}

func ForGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}