package main

import "fmt"

func main() {

	var N int
	var numbers []int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		numbers = append(numbers, 0)
		fmt.Scan(&numbers[i])
	}

	var max int
	var maxEnd int

	for i := 0; i < len(numbers); i++ {
		maxEnd += numbers[i]
		if max < maxEnd {
			max = maxEnd
		}
		if maxEnd < 0 {
			maxEnd = 0
		}
	}
	fmt.Println(max)

}
