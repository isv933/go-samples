package main

import "fmt"

func chunk(a []int, chunkSize int) [][]int {
	var result [][]int
	result = make([][]int, 1)
	for i, chunk := 0, 0; i < len(a); i++ {
		if len(result[chunk]) >= chunkSize {
			chunk++
			result = append(result, []int{})
		}
		result[chunk] = append(result[chunk], a[i])
	}

	return result
}

func main() {
	fmt.Println(chunk([]int{1, 2, 3, 4}, 3))

}
