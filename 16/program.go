package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	var n int
	_, _ = fmt.Fscanf(input, "%d\n", &n)

	input.Close()

	result := stairs(n)

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d\n", result))
	output.Close()
}

func stairs(n int) int {
	var i, out int

	if n <= 2 {
		return 1
	}

	out = 0
	for i = 0; i*2 <= n; i++ {
		out += stairs(i)

		if i*2 == n {
			out--
		}
	}
	return out
}
