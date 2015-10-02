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
	values := strings.Split(lines[0], " ")
	a, _ := strconv.Atoi(strings.Trim(values[0], "\t \r"))
	b, _ := strconv.Atoi(strings.Trim(values[1], "\t \r \n"))
	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d", a+b)), 0)
}
