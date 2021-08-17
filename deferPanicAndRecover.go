package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


// ---
// DEFER
// defer used to close resources
// defer are executed after exiting from scope
// ----

func deferPanicAndRecoverTest() {

	// ---
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")


	// ---
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
	fmt.Println()

	// ---
	a := "start"
	defer fmt.Println(a)
	a = "end"
	fmt.Println()



	// ---
	//doSomePanic()


	// ----
	/*
		First functions, then defers and then panics, so no resource leaks!
	*/
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error catch should we recover? ===> : ", err)

			//panic(err) // Note: we can re-panic as well
		}
	}()

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("incoming request")
		w.Write([]byte("Hello Go!"))
	})
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println()

	// ----

}

func doSomePanic() {
	panic("something bad happened")
	w := 1
	o := 0
	fmt.Println(w / o)
}
