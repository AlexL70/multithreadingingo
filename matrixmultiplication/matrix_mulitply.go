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
	matrixA = [matrixSize][matrixSize]int{
		{3, 1, -4},
		{2, -3, 1},
		{5, -2, 0},
	}
	matrixB = [matrixSize][matrixSize]int{
		{1, -2, -1},
		{0, 5, 4},
		{-1, -2, 3},
	}
	result = [matrixSize][matrixSize]int{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workoutRow(row int) {
	for column := 0; column < matrixSize; column++ {
		for i := 0; i < matrixSize; i++ {
			result[row][column] += matrixA[row][i] * matrixB[i][column]
		}
	}
}

func main() {
	fmt.Println("Working...")
	start := time.Now()
	for count := 0; count < 100; count++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		for row := 0; row < matrixSize; row++ {
			workoutRow(row)
			//fmt.Println(result[row])
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)
}
