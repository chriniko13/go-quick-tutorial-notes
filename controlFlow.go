package main

import "fmt"

func controlFlowTest() {


	// ---
	var readFromDisk bool = true
	if readFromDisk {
		fmt.Println("read from disk")
	}
	fmt.Println()


	// ----
	recordsById := map[string]student{
		"foo": {name: "foo"},
	}

	if pop, ok := recordsById["foo"]; ok {
		fmt.Println("pop: ", pop)
	}
	fmt.Println()


	// ----
	stud := returnStud()
	fmt.Println(stud)
	fmt.Println(*stud)


	// ----
	switch 4 {
	case 1, 5, 10:
		fmt.Println("first")
	case 2,4,6:
		fmt.Println("second")
	default:
		fmt.Println("default")
	}

	// ---
	i := 2
	switch {
	case i > 1:
		fmt.Println("i > 1")
		fallthrough
	case i < 1:
		fmt.Println("i < 1")
	default:
		fmt.Println("i == 1")
	}


	// ----
	var k interface{} = 1
	switch myVar := k.(type) {
	case int:
		fmt.Println("int 1 ", myVar)
		break
		fmt.Println("int 2")

	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")

	}

}

type student struct {
	name string
}

func returnStud() *student {
	return &student{
		name: "fun name",
	}
}