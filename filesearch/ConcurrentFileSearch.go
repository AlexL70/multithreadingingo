package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	rootDir   = flag.String("dir", "", "Directory where to start search (root)")
	fileName  = flag.String("file", "", "Name of file to search")
	matches   []string
	waitGroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, fileName string) {
	fmt.Println("Searching in", root)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.Name() == fileName {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), fileName)
		}
	}
	waitGroup.Done()
}

func main() {
	flag.Parse()
	if *rootDir == "" || *fileName == "" {
		fmt.Println("Some of parameters missed")
		flag.PrintDefaults()
		return
	}
	waitGroup.Add(1)
	go fileSearch(*rootDir, *fileName)
	waitGroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
}
