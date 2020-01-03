package slices

import "fmt"

func ManipulateSlices() {
	var hardware[] string

	hardware = append(hardware, "Placa de Video")
	fmt.Println(hardware)

	hardware = append(hardware, "Teclado", "Mouse")
	fmt.Println(hardware[1], hardware[2])

	fmt.Println(len(hardware))
}