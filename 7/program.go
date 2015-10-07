package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	var a, b, c string
	_, _ = fmt.Fscanf(input, "%s %s %s\n", &a, &b, &c)
	input.Close()

	output, _ := os.Create("output.txt")
	var result = max(max(a, b), c)
	output.WriteString(result + "\n")
	output.Close()

}

func max(a, b string) string {
	la := len(a)
	lb := len(b)

	if la > lb {
		return a
	} else {
		return b
	}

	if la == lb {
		for i := 0; i < la; i++ {
			if a[i] > b[i] {
				return a
			}
		}
	}
	return b
}
