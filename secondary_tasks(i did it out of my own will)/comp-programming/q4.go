package main

import (
	"fmt"
)

func ship(bottles int) {
	sizes := []int{48, 24, 12, 6}
	names := []string{"xl", "large", "medium", "small"}

	for i := 0; i < len(sizes); i++ {
		count := bottles / sizes[i]
		if count > 0 {
			fmt.Printf("%d %s ", count, names[i])
		}
		bottles %= sizes[i]
	}
	fmt.Println()
}

func main() {
	var bottles int
	fmt.Println("enter number of bottles to be shipped:")
	fmt.Scan(&bottles)

	ship(bottles)
}
