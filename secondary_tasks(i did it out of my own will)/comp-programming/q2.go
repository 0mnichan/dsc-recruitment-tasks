//2.  Remove duplicates from array

package main

import (
	"fmt"
)

func gorlockthedestroyer(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	result := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	var n int
	fmt.Println("enter number of elements in the array")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("enter elements of the array")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(gorlockthedestroyer(arr))
}
