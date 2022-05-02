package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy sees balance of", money)
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy done.")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		fmt.Println("Spendy sees balance of", money)
		lock.Unlock()
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
