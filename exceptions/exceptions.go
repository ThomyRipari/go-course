package exceptions

import (
	"fmt"
	"errors"
)

func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
		case string:
			return 0, errors.New("El valor A es un string")
	}

	switch b.(type) {
		case string:
			return 0, errors.New("El valor B es un string")
	}

	return a.(int) + b.(int), nil
}

func ManipulateErrors() {
	result, error := Sum(20, 23)

	if error != nil {
		panic(error)
	}

	fmt.Println(result, error)
}