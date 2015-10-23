package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	// when the mutex (m) is locked any other attempt to lock it
	// will block until it is unlocked. IT is recommended that
	// you use go's concurrency features, though.

	var input string
	fmt.Scanln(&input)
}
