package main

import "fmt"

func arraysAndSlices() {

	// ---
	grades := [...]int{8, 7, 9}
	fmt.Println(grades)
	fmt.Println(grades[0])
	fmt.Println(grades[0:3])


	grades[0] = 10
	fmt.Println(grades)
	fmt.Println(len(grades))



	// ---
	var mat = [3][3]int {{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}
	fmt.Println(mat)
	fmt.Println(mat[0])
	fmt.Println(mat[0][1])


	// ---
	a := [...]int{1,2,3}
	b:= a
	b[1] = 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	c := &a
	c[1] = 3
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println()


	// --- slice ---
	d := []int{1,2,3}
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	g := append(d, 4)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))
	fmt.Println()


	// ---
	l := make([]int, 3, 9)
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))
	fmt.Println()


	// ----
	p := []int{1,2,3,4,5}
	w := append(p[:2], a[3:]...)
	fmt.Println(w)


}
