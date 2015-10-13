package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	result := "NO\n"

	var a, b int
	_, _ = fmt.Fscanf(input, "%d %d\n", &a, &b)

	input.Close()

	
	output, _ := os.Create("output.txt")
	output.WriteString(result)
	output.Close()
}
