package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

	Creating goroutines

	Synchronization
		WaitGroups
		Mutexes

	Parallelism

	Best practises


	Check race conditions during compile time ---> go run -race

 */

var wg = sync.WaitGroup{}

var counter = 0


var m = sync.RWMutex{}

func goroutinesTest() {

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// ---
	wg.Add(1)
	go sayHello() // Note: spin and run this function in green thread (vs OS thread).


	// ---
	wg.Add(1)
	var msg = "hello there"
	go func(input string) {
		fmt.Println(input)

		m.Lock()
		counter++
		m.Unlock()

		wg.Done()
	}(msg)
	msg = "goodbye"



	wg.Wait()
	fmt.Println("counter is: ", counter)

	//time.Sleep(time.Second * 2)

}



func sayHello() {
	fmt.Println("Hello")

	m.Lock()
	counter++
	m.Unlock()

	wg.Done()
}
