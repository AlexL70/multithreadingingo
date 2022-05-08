package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Error opening page %s: %w", url, err)
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error reading body. Url: %s. Error: %w", url, err)
		fmt.Println(err)
		return
	}
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			atomic.AddInt32(&frequency[index], 1)
		}
	}
}

func main() {
	var frequency [26]int32
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 1000; i < 1200; i++ {
		wg.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
