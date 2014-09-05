package main

import (
	"fmt"
	"sync"
)

// start 1 OMIT
func main() {
	in := make(chan int)
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		select { // いずれか受信するまでブロック
		case v := <-in:
			fmt.Println(v)
		case out<-"out!!":
		}
		wg.Done()
	}()
	fmt.Println(<-out)
	wg.Wait()
}
// end 1 OMIT
