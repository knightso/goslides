package main

import "fmt"
import "time"

// start 1 OMIT
func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 送信！
			fmt.Printf("sent %d\n", i)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println(<-ch) // 受信！
		fmt.Printf("received %d\n", i)
	}
}

// end 1 OMIT
