package main

import (
	"fmt"
	"time"
)

func main() {
	// start 1 OMIT
	req := make(chan int)
	count := make(chan int)

	go func() { // メインループ
		var c int
		for {
			select {
			case v := <-req:
				fmt.Printf("received %d\n", v)
				c++
			case count <- c:
			}
		}
	}()
	// end 1 OMIT

	// start 2 OMIT
	go func() { // リクエスト送信ループ
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			req <- i
		}
	}()

	for { // 5秒に一回count確認
		time.Sleep(5 * time.Second)
		fmt.Printf("count=%d\n", <-count)
	}
	// end 2 OMIT
}

