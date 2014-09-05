package main

import (
	"fmt"
	"time"
)

/**
 goroutineとchannelでpromiseっぽく書いてみたけど、よく見たらgoroutine一個で十分だこれ。。
 てな訳でボツ。
*/
// start 1 OMIT
func someLongJob(i int) int {
	time.Sleep(time.Duration(i) * time.Second) // some long job
	return i + 1
}
func main() {

	promise1 := make(chan int)
	go func() {
		promise1 <- someLongJob(1)
	}()

	promise2 := make(chan int)
	go func() {
		promise2 <- someLongJob(<-promise1)
	}()

	promise3 := make(chan int)
	go func() {
		promise3 <- someLongJob(<-promise2)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(<-promise3)
}

// end 1 OMIT
