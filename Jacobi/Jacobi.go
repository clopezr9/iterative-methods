package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Array []float64

func main() {
	args := os.Args
	var N int = 100
	for i := 1; i < len(args); i++ {
		file_name := args[i]
		fmt.Println(file_name, ":")
		var a, b = readCSVFile(file_name)
		jacobi(a, b, N)
		//gaussSimple(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
	}

	//var A = [][]float64{{5.0, 2.0, -3}, {2.0, 10.0, -8.0}, {3.0, 8.0, 13.0}}
	//var b = []float64{1.0, 4.0, 7.0}
	//var N int = 100
}

func readCSVFile(file_name string) ([]Array, []float64) {
	csvFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	matrix := []Array{}
	b := []float64{}
	for _, line := range csvLines {
		line_length := len(line)
		array_aux := []float64{}
		for i := 0; i < line_length; i++ {
			line_i, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				fmt.Println(err)
			}
			if i == line_length-1 {
				b = append(b, line_i)
				break
			}
			array_aux = append(array_aux, line_i)
		}
		matrix = append(matrix, array_aux)
	}

	return matrix, b
}

func dominant(A []Array) bool {

	for i := 0; i < len(A); i++ {
		var dom = math.Abs(A[i][i])
		var count = 0.0
		for j := 0; j < len(A); j++ {
			count = count + math.Abs(A[i][j])
		}
		count = count - dom
		if dom < count {
			return false
		}
	}
	return true
}

func dot(A []Array, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		var aux = 0.0
		for j := 0; j < l; j++ {
			aux = aux + b[j]*A[i][j]
		}
		result[i] = aux
	}
	return result
}

func substraction(a []float64, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		result[i] = a[i] - b[i]
	}
	return result
}
func division(a []float64, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		result[i] = a[i] / b[i]
	}
	return result
}

func matrix_permutations(arr [][]float64) [][][]float64 {
	var helper func([][]float64, int)
	res := [][][]float64{}

	helper = func(arr [][]float64, n int) {
		if n == 1 {
			tmp := make([][]float64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutations(arr []float64) []Array {

	var helper func([]float64, int)
	res := []Array{}

	helper = func(arr []float64, n int) {
		if n == 1 {
			tmp := make([]float64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1.0]
					arr[n-1.0] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1.0]
					arr[n-1.0] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func jacobi(A []Array, b []float64, N int) {

	//Creates varible with the length of Matrix A
	var l int = len(A)
	var ig = make([]float64, l)

	//Variable were dominant values are going to be stored.
	D := make([]float64, l)
	for i := 0; i < len(A); i++ {
		//Create a vector of the diagonal elements of A.
		D[i] = A[i][i]
	}

	//Creating Matrix were result from A - D is going to be stored.
	R := make([]Array, l)
	for i := 0; i < len(A); i++ {
		R[i] = make([]float64, l)
	}

	// Filling R with correct values.
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i != j {
				R[i][j] = A[i][j]
			}
		}
	}

	start := time.Now()
	// Iterate for N times.
	for i := 0; i < N; i++ {
		ig = division(substraction(b, dot(R, ig, l), l), D, l)
	}
	duration := time.Since(start)
	fmt.Println("EXECUTION TIME: ", duration.Seconds())
	fmt.Println("Ans:", ig)
}
