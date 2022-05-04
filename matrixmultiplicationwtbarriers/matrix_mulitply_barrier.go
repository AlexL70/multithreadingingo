package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	result       = [matrixSize][matrixSize]int{}
	workStart    = NewBarrier(matrixSize + 1)
	workComplete = NewBarrier(matrixSize + 1)
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workoutRow(row int) {
	for {
		workStart.Wait()
		for column := 0; column < matrixSize; column++ {
			for i := 0; i < matrixSize; i++ {
				result[row][column] += matrixA[row][i] * matrixB[i][column]
			}
		}
		workComplete.Wait()
	}
}

func main() {
	fmt.Println("Working...")
	for row := 0; row < matrixSize; row++ {
		go workoutRow(row)
	}
	start := time.Now()
	for count := 0; count < 100; count++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		workStart.Wait()
		workComplete.Wait()
	}
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)
}
