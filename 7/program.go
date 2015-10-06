package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	numbers := strings.Split(strings.Trim(lines[0], "\t \r"), " ")
	var output = max(max(numbers[0], numbers[1]), numbers[2])
	ioutil.WriteFile("output.txt", []byte(output), 0)
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
