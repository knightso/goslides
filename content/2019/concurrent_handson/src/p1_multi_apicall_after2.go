package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("start")
	start := time.Now()

	// START OMIT
	var res1, res2, res3 int

	eg := new(errgroup.Group)

	eg.Go(func() (err error) {
		res1, err = callAPI()
		return err
	})

	eg.Go(func() (err error) {
		res2, err = callAPI()
		return err
	})

	eg.Go(func() (err error) {
		res3, err = callAPI()
		return err
	})

	if err := eg.Wait(); err != nil {
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
