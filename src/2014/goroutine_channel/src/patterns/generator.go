package main

import (
	"fmt"
)

// start 1 OMIT
func seq(max int) <-chan int {
	gen := make(chan int)
	go func() {
		for i := 0; i <= max; i++ {
			gen<-i
		}
		close(gen)
	}()
	return gen
}
func main() {
	g := seq(20)

	fmt.Println(<-g)
	fmt.Println(<-g)

	for i := range g { // rangeでイテレートも可能
		fmt.Println(i)
	}
}

// end 1 OMIT
