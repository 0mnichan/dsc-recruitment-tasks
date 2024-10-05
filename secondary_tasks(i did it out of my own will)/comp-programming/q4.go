//4. The bottle shipping problem

package main

import "fmt"

func calc(bottles int) {
	xl := bottles / 48
	bottles %= 48

	large := bottles / 24
	bottles %= 24

	medium := bottles / 12
	bottles %= 12

	small := bottles / 6
	bottles %= 6

	if bottles > 0 {
		small++
	}
	fmt.Printf("%d xl, %d large, %d medium, %d small\n", xl, large, medium, small)
}

func main() {
	var bottles int
	fmt.Print("enter number of bottles")
	fmt.Scan(&bottles)
	calc(bottles)
}
