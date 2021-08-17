package main

import (
	"fmt"
)

func primitivesTest() {

	// ---
	var n bool = true
	fmt.Printf("%v %T\n", n, n)


	// ---
	fmt.Println(123 | 121)

	// ---
	var scien float64 = 13.7e72
	fmt.Println(scien)

	// ---
	var clex complex64 = 1 + 3i
	fmt.Println(clex)
	fmt.Println(real(clex))
	fmt.Println(imag(clex))
	fmt.Println(complex(2, -4))


	// ---
	var s = "this is a string"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(string(s[2]))
	//s[2] = "u" // Note: strings are immutable
	b := []byte(s)
	fmt.Println(b)

	// ---
	var r rune = 'a'
	fmt.Println(r)
	fmt.Println(string(r))

	// ---
	whatIs(leaf{value: 12})

	var tree = node{
		left: node{left: leaf{value: 12},
			       right: node{
						left:  leaf{value: 11},
						right: leaf{value: 14}},
		      },
		      right: leaf{value: 1}}

	fmt.Println(tree)
	printValuesOfTree(tree)


}



type tree interface {
}

func whatIs(t tree) {

	switch t.(type) {
	case node:
		fmt.Println("is a node")
		break
	case leaf:
		fmt.Println("is a leaf")
		break
	}
}

func printValuesOfTree(t tree) {

	switch v := t.(type) {
	case node:
		printValuesOfTree(v.left)
		printValuesOfTree(v.right)
		break

	case leaf:
		fmt.Printf("value is %d\n", v.value)
		break
	}

}


type node struct {
	tree
	left  tree
	right tree
}

type leaf struct {
	tree
	value int
}
