package main

import (
	"fmt"
)

func main() {
	var numbers = []int{1, 2, 3, 4}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)
}
