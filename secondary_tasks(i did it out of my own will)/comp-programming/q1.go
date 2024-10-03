//1. Arrange words in a string in the order of their length

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sorter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a cool sentence")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	words := strings.Fields(sentence)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	fmt.Println(words)
}

func main() {
	sorter()
}
