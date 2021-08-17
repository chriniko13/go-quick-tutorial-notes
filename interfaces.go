package main

import (
	"bytes"
	"fmt"
	"log"
)

/*

	AGENDA:
		basics

		composting interfaces

		type conversions
			the empty interface
			type switches

		implementing with values vs pointers



	BEST PRACTISES:
		use many, small interfaces

		don't export interfaces for types that will be consumed

		do export interfaces for types that will be used by package




 */
func interfacesTest() {

	// ---
	var w Writer = ConsoleWriter{}
	_, err := w.Write([]byte("hello go!"))
	if err != nil {
		panic(err)
	}
	fmt.Println()



	// ---
	var globalIncrementer IntCounter = IntCounter(0)

	globalIncrementer.Increment()
	globalIncrementer.Increment()
	globalIncrementer.Increment()
	fmt.Println(globalIncrementer.Increment())
	fmt.Println()


	// ---
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error catch should we recover? ===> : ", err)
			//panic(err) // Note: we can re-panic as well
		}
	}()

	var writerCloser = BufferedWriterCloser{buffer: bytes.NewBuffer([]byte{})}
	_, err2 := writerCloser.Write([]byte("hello there chriniko"))
	if err2 != nil {
		panic(err2)
	}
	err3 := writerCloser.Close()
	if err3 != nil {
		panic(err3)
	}
	fmt.Println()


	// ---- empty interface -----
	var emptyInterface interface{}  = BufferedWriterCloser{buffer: bytes.NewBuffer([]byte{})}
	newVar, ok := emptyInterface.(BufferedWriterCloser)
	if ok {
		newVar.Write([]byte("hello"))
		newVar.Close()
	}


	var anInteger interface{} = 0
	switch anInteger.(type) {
	case int:
		fmt.Println("it is an int")
	default:
		fmt.Println("i dont know")
	}

}


// ---
type Writer interface {
	Write([]byte) (int, error)
}


type ConsoleWriter struct {}

func (c ConsoleWriter)Write(data []byte) (int, error) {
		n, err := fmt.Println(string(data))
		return n, err
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	var v []byte = make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}


func (bwc *BufferedWriterCloser)Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}


// ---
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
