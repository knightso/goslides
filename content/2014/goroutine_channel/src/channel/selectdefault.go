package main

import (
	"fmt"
	"sync"
)

// start 1 OMIT
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		default: // ch1, ch2のどちらも受信出来なかった場合
			fmt.Println("default!")
		}
		fmt.Println("done!")
		wg.Done()
	}()

	wg.Wait()
}
// end 1 OMIT
