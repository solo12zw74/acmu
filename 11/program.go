package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	var k, n int
	_, _ = fmt.Fscanf(input, "%d %d\n", &k, &n)

	input.Close()

	arr := make([]*big.Int, 301)

	arr[0] = big.NewInt(1)
	for i := 1; i <= n; i++ {
		arr[i] = big.NewInt(0)
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				arr[i].Add(arr[i], arr[i-j])
			}
		}
	}

	output, _ := os.Create("output.txt")
	output.WriteString(fmt.Sprintf("%d\n", arr[n]))
	output.Close()
}
