package main

import (
	"fmt"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int) (int, [][]int) {
	var result [][]int
	var nextIndex int

	for idx, number := range numbers {
		nextIndex = idx + 1
		
		for i := nextIndex; i < len(numbers); i++ {
			if nextIndex >= len(numbers) {
				break;
			}
			nextNumber := numbers[i]
			if number + nextNumber == input {
				var couple = []int{number, nextNumber}
				result = append(result, couple)
			}

		}

	}
	return input, result
}
