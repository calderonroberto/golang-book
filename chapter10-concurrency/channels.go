package main

import (
	"fmt"
	"time"
)

//func pinger(c chan<- string) //restrict to receiving
func pinger(c chan string) { //otherwise bidirectional
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//func ponger(c <-chan string) //restrict to sending
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		fmt.Println(<-c) //block until receiving message
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
