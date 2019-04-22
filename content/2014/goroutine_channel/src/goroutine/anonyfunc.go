package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hoge!")
	}()

	time.Sleep(time.Microsecond)
}
