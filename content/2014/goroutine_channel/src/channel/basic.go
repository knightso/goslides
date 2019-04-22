package main

import "fmt"

// start 1 OMIT
func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello!" // 送信！
		ch <- "world!" // 送信！
	}()

	fmt.Println(<-ch) // 受信！
	fmt.Println(<-ch) // 受信！
}
// end 1 OMIT
