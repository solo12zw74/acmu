package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")

	result := ""

	var a, b, c, d int
	_, _ = fmt.Fscanf(input, "%d %d %d %d\n", &a, &b, &c, &d)

	input.Close()

	for i := -100; i < 101; i++ {
		if a*i*i*i+b*i*i+c*i+d == 0 {
			result += fmt.Sprintf("%d ", i)
		}
	}

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%s\n", strings.Trim(result, " ")))
	output.Close()
}
