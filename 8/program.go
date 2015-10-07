package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	result := "NO\n"

	var a, b, c int
	_, _ = fmt.Fscanf(input, "%d %d %d\n", &a, &b, &c)

	input.Close()

	if a*b == c {
		result = "YES\n"
	}

	output, _ := os.Create("output.txt")
	output.WriteString(result)
	output.Close()
}
