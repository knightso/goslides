package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/knightso/base/errors"
)

func main() {
	fmt.Println("start")
	start := time.Now()

	// START1 OMIT
	urlsScanner := bufio.NewScanner(newURLsReader())

	urlCh := make(chan string)

	go func() {
		for urlsScanner.Scan() {
			line := urlsScanner.Text()
			urlCh <- line
		}
		if err := urlsScanner.Err(); err != nil {
			log.Fatal(err)
		}
		close(urlCh)
	}()

	linksCh := make(chan []string)

	go func() {
		wg := new(sync.WaitGroup)
		for url := range urlCh {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				links, err := findLinks(url)
				if err != nil {
					log.Fatal(err)
				}
				linksCh <- links
			}(url)
		}
		wg.Wait()
		close(linksCh)
	}()

	pr1, pw1 := io.Pipe()

	go func() {
		defer pw1.Close()

		for links := range linksCh {
			for _, link := range links {
				fmt.Println(link) // just log
				if _, err := fmt.Fprintln(pw1, link); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	pr2, pw2 := io.Pipe()

	go func() {
		defer pw2.Close()

		if err := writeZip("test.zip", pw2, pr1); err != nil {
			log.Fatal(err)
		}
	}()

	if err := writeToRemote(pr2); err != nil {
		log.Fatal(err)
	}

	// END1 OMIT

	fmt.Println("time: ", time.Now().Sub(start))
}

func newURLsReader() io.Reader {
	// mocking remote storage file
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		for i := 0; i < 1000; i++ {
			if _, err := fmt.Fprintf(w, "http://www.example.com/%d\n", i); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Millisecond) // mock latency
		}
	}()
	return r
}

func findLinks(url string) ([]string, error) {
	// mocking remote fetch
	latency := rand.Intn(10) + 1
	time.Sleep(time.Duration(latency) * time.Millisecond)

	links := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		links = append(links, fmt.Sprintf("%s/%d", url, i))
	}

	return links, nil
}

func writeZip(fileName string, w io.Writer, r io.Reader) error {
	zipWriter := zip.NewWriter(w)

	fileWriter, err := zipWriter.Create(fileName)
	if err != nil {
		return errors.WrapOr(err)
	}

	if count, err := io.Copy(fileWriter, r); err != nil {
		return errors.WrapOr(err)
	} else {
		fmt.Println("zipped", count)
	}

	if err := zipWriter.Flush(); err != nil {
		return errors.WrapOr(err)
	}

	if err := zipWriter.Close(); err != nil {
		return errors.WrapOr(err)
	}

	return nil
}

func writeToRemote(reader io.Reader) error {
	// mocking write to remote storage
	size, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		return err
	}

	fmt.Printf("wrote: %d\n", size)
	return nil
}
