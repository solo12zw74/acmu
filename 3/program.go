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
	line := lines[0]
	linelen := len(line)
	var n int = 0
	if linelen < 2 {
		ioutil.WriteFile("output.txt", []byte("25"), 0)
		return
	} else {
		n, _ = strconv.Atoi(line[:linelen-1])
	}

	result := n * (n + 1)

	ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d25", result)), 0)

}
