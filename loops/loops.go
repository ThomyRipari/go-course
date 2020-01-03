package loops

import "fmt"

func ManipulateLoops() {
	// traditional for

	for i := 0; i < 5; i++ {
		fmt.Println("[LOOP]", i)
	}

	// while

	j := 5

	for j > 0 {
		j -= 1

		fmt.Println("[WHILE]", j)
	}

}