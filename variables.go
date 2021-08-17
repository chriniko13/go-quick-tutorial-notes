package main

import "fmt"
import "strconv"

var actorName = "someone"

var (
	dog   = "Bruno"
	owner = "Nikos"
)

var i = 1992

var ACCESSIBLE = "access"

func variablesTest() {



	// ---
	fmt.Println("hello there!")

	// ---
	const name uint64 = 123
	fmt.Println(name)

	// ---
	fmt.Println(i)
	var i int
	var x int
	i = 42
	x = 17
	fmt.Println(i, " ", x)
	fmt.Println(actorName)

	// ---
	k := 99
	fmt.Printf("%v, %T\n", k, k)

	// ---
	fmt.Println(dog)


	// ---
	var j string
	j = strconv.Itoa(2)
	fmt.Println(j)


	// ---
	var intVar = 123
	var floatVar float32
	floatVar = float32(intVar)
	fmt.Println(floatVar)

	// ---
	type person struct {
		name string
		age  int
	}

	p := person{age: 12, name: "somename"}

	fmt.Println(p)
}
