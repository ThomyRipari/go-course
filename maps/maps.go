package maps

import "fmt"

func ManipulateMaps() map[string] int {
	persons := map[string] int {
		"Andres": 32,
		"Fabricio": 46,
		"Matias": 35,
	}
	persons["Ricardo"] = 34
	persons["Martin"] = 49

	fmt.Println(persons)

	return persons
}

func GetPersonName(name string) int {
	persons := ManipulateMaps()

	fmt.Println(persons[name])

	return persons[name]
} 

func DeletePerson(name string) map[string] int {
	persons := ManipulateMaps()

	delete(persons, name)

	fmt.Println(persons)

	return persons
}