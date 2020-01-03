package arrays

import "fmt"

func ManipulateArrays() {
	var frutas[3] string

	fmt.Println(frutas)
	fmt.Println(frutas[0], frutas[1], frutas[2])

	frutas[0] = "Pera"
	frutas[1] = "Banana"
	frutas[2] = "Manzana"

	fmt.Println(frutas)

	fideos := [3] string {"Spaghetti", "Mostachol", "Resortes"}

	fmt.Println(fideos)

	// using make function

	software := make([] string, 2)

	software[0] = "Windows"
	software[1] = "Ubuntu"

	fmt.Println(software)

	// two dimensions array

	numbers := make([][] int, 3)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = make([] int, 2)

		for j := 0; j < 2; j++ {
			numbers[i][j] = i + j
		}
	}

	fmt.Println(numbers)
}
