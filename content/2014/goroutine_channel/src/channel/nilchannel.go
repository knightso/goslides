package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	var ch2 chan int

	go func() {
		ch1<-999
	}()

	select {
	case i := <-ch1:
		fmt.Println(i)
	case i := <-ch2:
		fmt.Println(i)
	}
}
