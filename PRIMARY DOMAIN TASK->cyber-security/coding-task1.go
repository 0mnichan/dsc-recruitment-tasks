// check for at least one lowercase, uppercase, number, and symbol

package main

import "fmt"

func checkpwd() {
	var pwd string
	var lower, upper, digit, symbol bool

	fmt.Println("pls enter your pwd")
	fmt.Scanln(&pwd)

	for _, char := range pwd {
		if char >= 'a' && char <= 'z' {
			lower = true
		} else if char >= 'A' && char <= 'Z' {
			upper = true
		} else if char >= '0' && char <= '9' {
			digit = true
		} else {
			symbol = true
		}
	}

	if lower && upper && digit && symbol {
		fmt.Println("nice pwd! it has all the requirements")
	} else {
		fmt.Println("wrong! check the following:")
		if !lower {
			fmt.Println(" at least one lowercase letter")
		}
		if !upper {
			fmt.Println(" at least one uppercase letter")
		}
		if !digit {
			fmt.Println(" at least one digit ")
		}
		if !symbol {
			fmt.Println(" at least one symbol required")
		}
	}
}

func main() {
	checkpwd()
}
