package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
	channel basics

	restricting data flow

	buffered channels

	closing channels

	for...range loops with channels

	select statements
*/

var waitGroup = sync.WaitGroup{}

func channelsTest() {

	// ---
	ch := make(chan string) // unbuffered channel.

	waitGroup.Add(2)

	go func() {
		i := <-ch
		fmt.Println("received: ", i)
		waitGroup.Done()
	}()

	go func() {
		ch <- "message"
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("--------------------")
	fmt.Println()

	// ----
	achievements_chan := make(chan achievement_unlocked_message, 50) // buffered channel.
	coord := sync.WaitGroup{}

	coord.Add(2)
	go achievementsPublisher(achievements_chan, &coord)
	go achievementsConsumer(achievements_chan, &coord)

	coord.Wait()
	//defer func() {
	//	close(achievements_chan)
	//}()



	// ---
	var doneCh = make(chan struct {}) // signal channel

	go logger(achievements_chan, doneCh, &coord)
	doneCh <- struct{}{}


}

// ----
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func achievementsPublisher(out chan<- achievement_unlocked_message, coord *sync.WaitGroup) {

	for _, v := range pow {
		var msg = "achievement_unlocked_" + strconv.Itoa(v)
		out <- achievement_unlocked_message{badge: msg}
	}

	close(out)
	(*coord).Done()
}

func achievementsConsumer(in <-chan achievement_unlocked_message, coord *sync.WaitGroup) {

	for {
		if message_received, ok := <-in; ok {
			fmt.Println("message received: ", message_received)
		} else {
			break
		}

	//for message_received := range in {

	//for i := 0; i < len(pow); i++ {
	//	message_received := <-in

	//	fmt.Println("message received: ", message_received)

	}
	coord.Done()
}

type achievement_unlocked_message struct {
	badge string
}

// ---

func logger(messages <-chan achievement_unlocked_message, doneCh <-chan struct{}, coord *sync.WaitGroup) {

	for {
		select {
		case entry := <-messages:
			fmt.Println("entry consumed from logger: ", entry)

		case <-doneCh:
			fmt.Println("DONE!")
			break
		default: // Non-blocking.

		}
	}
	coord.Done()

}
