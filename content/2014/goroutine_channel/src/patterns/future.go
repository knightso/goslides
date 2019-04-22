package main

import (
	"fmt"
	"time"
)

// start 1 OMIT
func someLongJob(i int) <-chan string {
	future := make(chan string)
	go func() {
		time.Sleep(time.Duration(5 - i) * time.Second) // some long job
		future <- fmt.Sprintf("success! %d", i)
	}()
	return future
}
func main() {
	futures := [](<-chan string){}
	for i := 0; i < 5; i++ { // 先ずGoroutineを全て起動
		futures = append(futures, someLongJob(i))
	}
	for _, future := range futures { // 全て起動し終わってから結果を参照
		fmt.Println(<-future)
	}
}

// end 1 OMIT
