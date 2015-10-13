package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// select is like a switch for channels
			// often used to implement a timeout
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	// Channels can be buffered
	// c := make(chan int, 1)
	// this will create a channel with a capacity of 1.
	// normally channels are synchronous; both sides of
	// the channel will wait until the other side is ready.
	// a buffered channel is ASYNCRONOUS, sending or receiving
	// will not wait unless the channel is already full.

	var input string
	fmt.Scanln(&input)

}
