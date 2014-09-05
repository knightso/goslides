package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan<-int)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			for number := 1; number < 27; number++ {
				fmt.Printf("%d ", number)
				ch<-number
			}
		}
	}()

	for i := range (chan int)(ch) {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
