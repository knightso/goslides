package main

import (
	"fmt"
	"math/rand"
	"time"
)

type strOrErr struct {
	value string
	err   error
}

func main() {
	fmt.Println("start")
	start := time.Now()

	// START OMIT
	reslist := make([]strOrErr, 0, 100)

	for i := 0; i < 100; i++ {
		res, err := callAPI(i)
		reslist = append(reslist, strOrErr{res, err})
	}

	for _, res := range reslist {
		fmt.Printf("res:%s, err:%v\n", res.value, res.err)
	}
	// END OMIT

	fmt.Println("time: ", time.Now().Sub(start))
}

// STARTAPI OMIT
func callAPI(id int) (string, error) {
	// mocking IO latency
	latency := rand.Intn(100) + 1
	time.Sleep(time.Duration(latency) * time.Millisecond)

	// mocking error
	if latency%3 == 0 {
		return "", fmt.Errorf("got stupid! latency:%d", latency)
	}

	return fmt.Sprintf("%d:%d", id, latency), nil
}

// ENDAPI OMIT
