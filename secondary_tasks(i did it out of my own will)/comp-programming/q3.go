//3. Print the date based on the entered day and year

package main

import (
	"fmt"
)

func date(day int, year int) string {
	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		months[1] = 29
	}

	monthnames := []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}

	month := 0
	for i := 0; i < len(months); i++ {
		if day <= months[i] {
			month = i
			break
		}
		day -= months[i]
	}

	date := fmt.Sprintf("%d %s %d", day, monthnames[month], year)
	return date
}

func main() {
	var day, year int
	fmt.Println("enter the day number and the year")
	fmt.Scan(&day, &year)

	result := date(day, year)
	fmt.Println(result)
}
