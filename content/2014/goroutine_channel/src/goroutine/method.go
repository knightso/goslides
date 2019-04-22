package main

import (
	"fmt"
	"time"
)

// start 1 OMIT
type hoge string

func (h hoge) hello() {
	fmt.Printf("hello %s!\n", h)
}

func main() {
	var h hoge = "hoge"
	go h.hello()
	time.Sleep(time.Microsecond)
}
// end 1 OMIT
