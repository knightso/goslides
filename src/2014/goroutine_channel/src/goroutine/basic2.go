package main

import (
	"fmt"
	"time"
)

func hoge() {
	fmt.Println("hoge!")
}

func main() {
	go hoge()
	time.Sleep(time.Microsecond)
}
