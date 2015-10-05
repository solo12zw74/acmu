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
		num, _ := strconv.Atoi(strings.Trim(stringsArray[i], "\t \r"))
		if num%2 == 0 {
			even = append(even, fmt.Sprintf("%d", num))
		} else {
			odd = append(odd, fmt.Sprintf("%d", num))
		}
	}
	var result string = "YES"
	if len(odd) > len(even) {
		result = "NO"
	}

	output := strings.Replace(fmt.Sprintf("%v\n%v\n%s\n", odd, even, result), "[", "", 2)
	output = strings.Replace(output, "]", "", 2)

	ioutil.WriteFile("output.txt", []byte(output), 0)
}
