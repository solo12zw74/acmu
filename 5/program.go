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
	length, _ := strconv.Atoi(strings.Trim(lines[0], "\t \r"))
	stringsArray := strings.Split(lines[1], " ")

	var even []string
	var odd []string

	for i := 0; i < length; i++ {
		num := strconv.Atoi(strings.Trim(stringsArray[i], "\t \r"))
		if num%2 == 0 {
			even = add(even, fmt.Sprintf("%d", ))
		} else {

		}
	}

	var result int
	for i := 0; i <= n; i++ {
		result += i
	}
	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d", result)), 0)
}
