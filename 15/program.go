package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	dim, _ := strconv.Atoi(lines[0])
	matrix := make([][]int, dim)

	for i := 1; i < dim+1; i++ {
		numbers := strings.Split(lines[i], " ")
		matrix[i-1] = make([]int, dim)
		for j := 0; j < dim; j++ {			
			matrix[i-1][j], _ = strconv.Atoi(numbers[j])
		}
	}

	result := 0	
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if i == j || matrix[i][j] == 0 {
				continue
			} 
			
			if matrix[i][j] == 1 && matrix[j][i] == 1 {
				result++
			}
		}
	}

	result = result/2

	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d\n", result)), 0)

}
