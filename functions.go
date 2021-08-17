package main

import "fmt"

/*
	Functions and methods
 */
func functionsTest() {

	// ---
	func() {
		fmt.Println("anonymous")
	}()

	var myFunc func() = func() { fmt.Println("my function as var") }
	myFunc()

	// ---
	msg := "hey there function"
	sayMessage(msg)
	fmt.Println(msg)
	fmt.Println()

	sayMessageRef(&msg)
	fmt.Println(msg)
	fmt.Println()

	fmt.Println(*sum(1, 2, 3, 4))
	fmt.Println(sum2(1, 2, 3, 4))
	fmt.Println()

	// ----
	d, err := divide(1, 0)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("division: ", d)
	}
	fmt.Println()


	// ----
	p := myStruct{name: "Nick"}
	p.greet()


}

// ----
func sayMessage(msg string) {
	fmt.Println(msg)
	msg = msg + " changed"
}

func sayMessageRef(msg *string) {
	fmt.Println(*msg)
	*msg = *msg + " changed"
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sum2(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		//panic("cannot provide zero as second value")
		return 0.0, fmt.Errorf("cannot provide zero as second value")
	}
	return a / b, nil
}


// ----
type myStruct struct {
	name string
}
func (m *myStruct) greet() {
	fmt.Println("my name is: ", m.name)
}