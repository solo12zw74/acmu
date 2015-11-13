package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	var a, b string
	_, _ = fmt.Fscanf(input, "%s %s\n", &a, &b)

	input.Close()

	y, z := 0, 0

	for i := 0; i < 4; i++ {
		if a[i] == b[i] {
			y++
		}

		for j := 0; j < 4; j++ {
			if a[i] == b[j] && i != j {
				z++
			}
		}
	}

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d %d\n", y, z))
	output.Close()
}
