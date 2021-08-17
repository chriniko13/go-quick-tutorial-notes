package main

import "fmt"

func loopingTest() {

	// ---
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}
	fmt.Println()


	// ---
	var k int = 0
	for {
		fmt.Println("hey hey my my....")
		k++
		if k == 5 {
			break
		}
	}


	// ---
	s := []int{1,2,3}
	for idx, v := range s {
		fmt.Printf("%d ==> %d\n", idx, v)
	}

}
