package main

import (
	"fmt"
	"reflect"
)

func mapsAndStructsTest() {

	// ---
	statePopulations := map[string]int{
		"California": 39250017,
		"Ohio":       11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])
	fmt.Println(statePopulations["Ohio2"])
	delete(statePopulations, "Ohio")
	fmt.Println(statePopulations)
	fmt.Println()

	pop, ok := statePopulations["Ohio2"]
	if (ok) {
		fmt.Println("ok ", pop)
	} else {
		fmt.Println("not okay")
	}
	fmt.Println()


	// ---
	myMap := map[[3]int]int{
		{1, 2, 3}: 1,
	}
	fmt.Println(myMap)
	fmt.Println(myMap[[3]int{1, 2, 3}])
	fmt.Println()


	// ---
	aDoctor := Doctor{
		number: 3,
		name: "some doc",
		patients: []string{"foo", "bar", "baz"},
	}
	fmt.Println(aDoctor.patients[2])

	anonymStruct := struct {
		name string
	}{name: "foo"}
	fmt.Println(anonymStruct)
	fmt.Println()


	// ---
	animal := Bird{
		Animal: Animal{
			Name: "animal",
			Origin: "foobar",
		},
		flies: true,
	}
	fmt.Println(animal)

	t:= reflect.TypeOf(Animal{})
	field, fieldFound := t.FieldByName("Name")
	if (fieldFound) {
		fmt.Println(field.Tag)
	}
}


type Doctor struct {
	number int
	name string
	patients []string
}


type Animal struct {
	Name string `required max: "100"`
	Origin string
}
type Bird struct {
	Animal
	flies bool
}