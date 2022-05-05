package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy done.")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy done.")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(money)
}
