package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	start := time.Now()

	// START OMIT
	var res1, res2, res3 int
	var err1, err2, err3 error

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		res1, err1 = callAPI()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res2, err2 = callAPI()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res3, err3 = callAPI()
	}()

	wg.Wait()

	if err1 != nil {
		log.Fatal(err1)
	}

	if err2 != nil {
		log.Fatal(err2)
	}

	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println(res1, res2, res3)
	// END OMIT

	fmt.Println("time: ", time.Now().Sub(start))
}

// STARTAPI OMIT
func callAPI() (int, error) {
	// mocking IO latency
	latency := rand.Intn(3) + 1
	time.Sleep(time.Duration(latency) * time.Second)

	return latency, nil
}

// ENDAPI OMIT
