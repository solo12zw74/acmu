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
	result := 9 - n
	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d9%d", n, result)), 0)
}
