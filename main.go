package main

import (
	"fmt"
	"time"
	"github.com/thomyripari/curso-go/loops"
	"github.com/thomyripari/curso-go/slices"
	"github.com/thomyripari/curso-go/arrays"
	"github.com/thomyripari/curso-go/strings"
	"github.com/thomyripari/curso-go/maps"
	"github.com/thomyripari/curso-go/structs"
	"github.com/thomyripari/curso-go/exceptions"
	"github.com/thomyripari/curso-go/pointers"
	"github.com/thomyripari/curso-go/routines"
)

func main() {
	fmt.Println("--- BIENVENIDO A MI CALCULADORA --- ")

	a, b := getNumbers()
	
	operation := selectOperation()

	var result int

	if operation == 1 {
		result = a + b
	} else if operation == 2 {
		result = a - b
	} else if operation == 3 {
		result = a * b
	} else {
		result = a / b
	}

	fmt.Printf("El resultado es: %d \n", result)

	strings.ManipulateStrings()
	arrays.ManipulateArrays()
	slices.ManipulateSlices()
	loops.ManipulateLoops()
	maps.GetPersonName("Fabricio")
	maps.DeletePerson("Andres")
	structs.ManipulateStructs()
	exceptions.ManipulateErrors()
	pointers.ManipulatePointers()


	go routines.ForGo(300)
	go routines.ForGo(400)

	time.Sleep(10 * time.Second)
}


func getNumbers() (int, int) {
	var numberA int 
	var numberB int 

	fmt.Print("Ingrese el primer numero: ") 
	fmt.Scanf("%d", &numberA)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanf("%d", &numberB)

	return numberA, numberB
}

func selectOperation() int {
	fmt.Println(" --- ELIGA LA OPERACION QUE DESEA REALIZAR ---")

	var operation int

	fmt.Println("1 - SUMA")
	fmt.Println("2 - RESTA")
	fmt.Println("3 - MULTIPLICACION")
	fmt.Println("4 - DIVISION")

	fmt.Print(": ")

	fmt.Scanf("%d", &operation)

	return operation
}