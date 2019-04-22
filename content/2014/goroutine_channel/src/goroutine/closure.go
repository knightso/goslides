package main

import (
	"fmt"
	"time"
)

// start 1 OMIT
func inc() func() {
	i := 0
	return func() {
		i++
		fmt.Printf("i=%d\n", i)
	}
}

func main() {
	f := inc()
	go f()
	go f()
	go f()
	time.Sleep(time.Second)
}
// end 1 OMIT
