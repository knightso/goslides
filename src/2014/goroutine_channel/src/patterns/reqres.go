package main

import (
	"fmt"
	"time"
	"errors"
)

func main() {
	// start 1 OMIT
	req := make(chan int)
	count := make(chan int)
	stop := make(chan chan error) // ここがchan chan

	go func() { // メインループ
		var c int
		for {
			select {
			case v := <-req:
				fmt.Printf("received %d\n", v)
				c++
			case count <- c:
			case sch := <- stop:
				// tried some stop process, but failed
				sch<-errors.New("stop failed...")
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

	go func() {
		<-time.After(15 * time.Second) // 15秒後に終了要求
		sch := make(chan error)
		stop<-sch
		if err := <-sch; err != nil {
			fmt.Println(err)
		}
	}()

	c := time.Tick(5 * time.Second) // 5秒に一回count確認
	for _ = range c {
		fmt.Printf("count=%d\n", <-count)
	}
	// end 2 OMIT
}

