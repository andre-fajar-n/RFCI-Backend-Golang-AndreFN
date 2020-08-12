package main

import (
	"fmt"
)

func swap(arr []int, count int) ([]int, bool) {
	var index int = 0
	
	for index < len(arr)-1 {
		if (arr[index] > arr[index + 1]) {
			arr[index], arr[index+1] = arr[index+1], arr[index]
			fmt.Printf("%d. [%d,%d] -> ", count, arr[index], arr[index+1])
			fmt.Println(arr)
			return arr, true
		}
		index++
	}
	return arr, false
}

func CountSwap(arr []int) int {
	var count int = 0

	var sortArray, status = swap(arr, count+1)

	if !status {
		return count
	} else {
		count++
		for status {
			sortArray, status = swap(sortArray, count+1)
			count++
		}
		return count - 1
	}

}

func main()  {
	var data = []int{4,9,7,5,8,9,3}
	// var data = []int{1,2,3,4}
	fmt.Printf("Jumlah swap: %d\n", CountSwap(data))
}