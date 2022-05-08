package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int) error {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Error opening page %s: %w", url, err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error reading body. Url: %s. Error: %w", url, err)
		return err
	}
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			frequency[index]++
		}
	}
	return nil
}

func main() {
	var frequency [26]int
	start := time.Now()
	for i := 1000; i < 1200; i++ {
		err := countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
		if err != nil {
			fmt.Println(err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
