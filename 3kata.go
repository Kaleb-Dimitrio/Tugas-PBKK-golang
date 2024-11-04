package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan 3 kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)
	if len(words) < 3 {
		fmt.Println("Kata kurang dari 3")
		return
	}

	var reversedWords []string
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, reverseString(words[i]))
	}

	output := strings.Join(reversedWords, " ")
	fmt.Println(output)
}
