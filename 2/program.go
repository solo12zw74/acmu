package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	n, _ := strconv.Atoi(strings.Trim(lines[0], "\t \r"))
	var result int
	for i := 0; i <= n; i++ {
		result += i
	}
	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d", result)), 0)
}
