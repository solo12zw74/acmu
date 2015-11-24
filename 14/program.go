package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	var a, b int
	_, _ = fmt.Fscanf(input, "%d %d\n", &a, &b)

	input.Close()

	NOD := getGCD(a,b)
	result := a * b / NOD
	
	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d\n", result))
	output.Close()
}

func getGCD(a,b int) int {
	if a == b {
		return a
	}

	gcd := 0

	for ;a * b > 0; {
		if a > b { 
     		a = a % b
		} else { 
 			b = b % a;
 		}
   		gcd= a + b;
	}
	return gcd
}
