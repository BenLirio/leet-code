package main

import (
	"sync"
	"fmt"
)

type Item struct {
	left int
	right int
	top int
	bottom int
	val int
}

func newItemMatrix(data [][]int) [][]Item {
	size := len(data)
	matrix := make([][]Item, size)
	var wg sync.WaitGroup
	wg.Add(size)
	for i, row := range matrix {
		go func(row *[]Item, dataRow []int) {
			defer wg.Done()
		}(&row, data[i])
	}
	wg.Wait()
	return matrix
}


func longestIncreasingPath(data [][]int) int {
	matrix := newItemMatrix(data)
	fmt.Println(matrix)
	return 0
}

func main() {
	data := [][]int{
		{9,9,4},
		{6,6,8},
		{2,1,1}}
	longestIncreasingPath(data)
	fmt.Println("MAIN")
}
