package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/knightso/base/errors"
)

func main() {
	fmt.Println("start")
	start := time.Now()

	// START1 OMIT
	urlsScanner := bufio.NewScanner(newURLsReader())

	buf := bytes.NewBuffer(make([]byte, 0, 2048))

	for urlsScanner.Scan() {
		line := urlsScanner.Text()

		links, err := findLinks(line)
		if err != nil {
			log.Fatal(err)
		}

		for _, link := range links {
			if _, err := fmt.Fprintln(buf, link); err != nil {
				log.Fatal(err)
			}
		}
	}
	time.Sleep(1000)

	if err := urlsScanner.Err(); err != nil {
		log.Fatal(err)
	}
	// END1 OMIT

	// START2 OMIT
	fmt.Println(buf.String()) // just log

	zipBuf := bytes.NewBuffer(make([]byte, 0, 2048))

	if err := writeZip("test.zip", zipBuf, buf); err != nil {
		log.Fatal(err)
	}

	if err := writeToRemote(zipBuf); err != nil {
		log.Fatal(err)
	}

	// END2 OMIT

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
