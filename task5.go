package main

import (
	"fmt"
	"math/rand"
)

type Matrix2d [][]int

func printMatrix(mat Matrix2d) {
	for _, col := range mat {
		for _, value := range col {
			fmt.Printf("%3d ", value)
		}
		fmt.Print("\n")
	}
}

/* Вирішити завдання за допомогою масиву або зрізу (за вашим вибором) */
/* Знайти суму парних чисел четвертого рядка матриці */
func main() {
	var matSize int
	var err error
	matMin, matMax := 0, 128
	// rand.Seed(42)

	fmt.Print("Matrix size (at least 4): ")
	_, err = fmt.Scan(&matSize)
	if err != nil {
		panic(err)
	}
	if matSize < 4 {
		panic("Matrix size should be at least 4!")
	}

	mat := make(Matrix2d, matSize)
	for i := 0; i < matSize; i++ {
		mat[i] = make([]int, matSize)
		for j := 0; j < matSize; j++ {
			mat[i][j] = rand.Intn(matMax-matMin) + matMin
		}
	}

	fmt.Println("Matrix:")
	printMatrix(mat)

	fourthRow := mat[3]
	result := 0

	for _, rowNum := range fourthRow {
		if (rowNum % 2) == 0 {
			result += rowNum
		}
	}

	fmt.Printf("Result: %d\n", result)
}
