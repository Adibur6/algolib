package main

import (
	"fmt"
	"github.com/adibur6/algolib/algorithms"
)

func main() {
	arr := []int{432, 23, 534, 12, 6757, 34, 2342, 546, 23, 6756345, 456}
	fmt.Println("Unsorted Array:", arr)
	sortedArr := algorithms.BubbleSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
