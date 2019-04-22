package main

import (
	"fmt"
	"time"
	"sync"
)

// start 1 OMIT
var balance int = 0
var mutex sync.Mutex
func deposit(money int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine終了を通知
	mutex.Lock()
	defer mutex.Unlock()
	newbalance := balance + money
	time.Sleep(time.Microsecond)
	balance = newbalance
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3) // ３個のGoroutineを待つよ！
	go deposit(100, &wg)
	go deposit(100, &wg)
	go deposit(100, &wg)
	wg.Wait() // 全部終了まで待機！
	fmt.Printf("balance:%d\n", balance)
}
// end 1 OMIT
