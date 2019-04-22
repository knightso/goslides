package main

import (
	"fmt"
	"time"
	"sync"
)

// start 1 OMIT
var balance int = 0
var mutex sync.Mutex
func deposit(money int) {
	mutex.Lock()
	defer mutex.Unlock()
	newbalance := balance + money
	time.Sleep(time.Microsecond)
	balance = newbalance
}
// end 1 OMIT
// start 2 OMIT
func main() {
	go deposit(100)
	go deposit(100)
	go deposit(100)
	time.Sleep(time.Second) // これ！
	fmt.Printf("balance:%d\n", balance)
}
// end 2 OMIT
