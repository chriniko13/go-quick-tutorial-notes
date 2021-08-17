package main

import "fmt"

/*
	Typed constants
	Untyped constants
	Enumerated constants
	Enumeration expressions
 */
const MyConst string = "hey there exported"

const (
	_ = iota
	a
	b
	c
)


func constantsTest() {

	// ---
	const myConst string = "hey there"
	//myConst = "cannot do this"
	fmt.Println(myConst)


	// ---
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)


	// ---
	var arr []string = []string{"a", "b"}
	for idx, v := range arr {
		fmt.Printf("idx: %d value: %v\n", idx, v)
	}
}
