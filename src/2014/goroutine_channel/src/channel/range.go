package main

import "fmt"

// start 1 OMIT
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 送信！
			fmt.Printf("sent %d\n", i)
		}
		close(ch) 
	}()

	for i := range ch{
		fmt.Printf("received %d\n", i)
	}
}
// end 1 OMIT
