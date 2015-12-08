package main

import (
	"fmt"
	"local/acmu/common"
	"os"
	"strings"
)

func main() {
	lines, _ := common.ReadLines("input.txt")

	str := strings.Replace(lines[1], " ", "", -1)

	l := len(str)

	result := 1

	for i := l - 1; i > 0; i-- {
		substr := str[:i]
		r := strings.LastIndex(str, substr)
		if r > 1 && str[i] != str[i+1] {
			result = r
			break
		}
	}

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d\n", result))
	output.Close()
}
