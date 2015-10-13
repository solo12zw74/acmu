package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line1 := scanner.Text()

	len, _ := strconv.Atoi(line1)
	scanner.Scan()
	line2 := scanner.Text()
	sarr := strings.Split(line2, " ")

	arr := make([]int, len)

	var min, max, sum, prod, minPos, maxPos int

	prod = 1

	for i := 0; i < len; i++ {
		val, _ := strconv.Atoi(sarr[i])
		arr[i] = val
		if val < min {
			min, minPos = val, i
		}
		if val >= max {
			max, maxPos = val, i
		}

		if val > 0 {
			sum += val
		}
	}

	if maxPos < minPos {
		t := maxPos
		maxPos = minPos
		minPos = t
	}

	for i := minPos + 1; i < maxPos; i++ {
		prod *= arr[i]
	}

	if maxPos-minPos < 2 {
		prod = 0
	}

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d %d\n", sum, prod))
	output.Close()
}
