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
		select { // いずれか受信するまでブロック
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		}
		wg.Done()
	}()

	ch2<-999
	wg.Wait()
}
// end 1 OMIT
