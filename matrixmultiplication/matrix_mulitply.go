package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA   = [matrixSize][matrixSize]int{}
	matrixB   = [matrixSize][matrixSize]int{}
	result    = [matrixSize][matrixSize]int{}
	rwLock    = sync.RWMutex{}
	cond      = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workoutRow(row int) {
	rwLock.RLock()
	for {
		waitGroup.Done()
		cond.Wait()
		for column := 0; column < matrixSize; column++ {
			for i := 0; i < matrixSize; i++ {
				result[row][column] += matrixA[row][i] * matrixB[i][column]
			}
		}
	}
}

func main() {
	fmt.Println("Working...")
	waitGroup.Add(matrixSize)
	for row := 0; row < matrixSize; row++ {
		go workoutRow(row)
	}
	start := time.Now()
	for count := 0; count < 100; count++ {
		waitGroup.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		waitGroup.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)
}
