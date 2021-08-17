package main

import "fmt"

/*
	Creating pointers

	De-referencing pointers

	The new function

	Working with nul

	Types with internal pointers
*/
func pointersTest() {

	// ---
	a := 42
	b := a
	fmt.Println(&a, " --- ", &b)
	fmt.Println()


	// ---
	var i *int = new(int)
	*i = 17
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(*i)


	// ---
	arr := [3]int{1,2,3}
	var d *int = &arr[0]
	fmt.Println(d)
	fmt.Println(*d)
	fmt.Println()

	// ---
	var dbR = new(dbRecord)
	(*dbR).contents = new(string)
	*(*dbR).contents = "some contents"
	fmt.Println(dbR)
	fmt.Println(*dbR)
	fmt.Println(*(*dbR).contents)
	*dbR.contents = "help from compiler"
	fmt.Println(*dbR.contents)
	fmt.Println()


	// ---slice vs array copy---
	arrk := [3]int{1,2,3}
	arrkb := arrk
	fmt.Println(arrk, arrkb)
	arrk[1] = 111
	fmt.Println(arrk, arrkb)


}

type dbRecord struct {
	contents *string
}
