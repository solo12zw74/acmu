package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	length, _ := strconv.Atoi(strings.Trim(lines[0], "\t \r"))

	var xp, yp, x1, y1, x2, y2, x3, y3, x4, y4 int

	result := 0

	for i := 1; i <= length; i++ {
		_, _ = fmt.Sscanf(strings.Trim(lines[i], "\t \r"), "%d %d %d %d %d %d %d %d %d %d", &xp, &yp, &x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)
		if isInRect(xp, yp, x1, y1, x2, y2, x3, y3, x4, y4) {
			result++
		}
	}
	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d", result))
	output.Close()
}

func isInRect(xp, yp, x1, y1, x2, y2, x3, y3, x4, y4 int) bool {
	D1 := (x1-x2)*(yp-y2) - (y1-y2)*(xp-x2)
	D2 := (x2-x3)*(yp-y3) - (y2-y3)*(xp-x3)
	D3 := (x3-x4)*(yp-y4) - (y3-y4)*(xp-x4)
	D4 := (x4-x1)*(yp-y1) - (y4-y1)*(xp-x1)

	if D1 < 0 || D2 < 0 || D3 < 0 || D4 < 0 {
		return false
	}

	return true
}
