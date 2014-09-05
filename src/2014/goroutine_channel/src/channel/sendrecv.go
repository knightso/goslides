package main

import "fmt"

// start 1 OMIT
func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // 送信！
	}
	close(ch)
}
func receive(ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("received %d\n", i)
	}
}
func main() {
	ch := make(chan int)

	go send(ch)

	receive(ch)
}
// end 1 OMIT
