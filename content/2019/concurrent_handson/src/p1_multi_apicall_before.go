package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("start")
	start := time.Now()

	// START OMIT
	res1, err := callAPI()
	if err != nil {
		log.Fatal(err)
	}

	res2, err := callAPI()
	if err != nil {
		log.Fatal(err)
	}

	res3, err := callAPI()
	if err != nil {
		log.Fatal(err)
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
