package main

import (
	"fmt"
)

func reverse(s string) string {
	intString := []rune(s)
	for left, right := 0, len(intString)-1; left < right; left, right = left+1, right-1 {
		intString[left], intString[right] = intString[right], intString[left]

		left++
		right--
	}

	return string(intString)
}

func main() {
	fmt.Println(reverse("привет"))

}
