package main

import (
	"fmt"
	"time"
)

// start 1 OMIT
var balance int = 0
func deposit(money int) {
	newbalance := balance + money
	balance = newbalance
}
func main() {
	go deposit(100)
	go deposit(100)
	go deposit(100)
	time.Sleep(time.Second)
	fmt.Printf("balance:%d\n", balance)
}
// end 1 OMIT
