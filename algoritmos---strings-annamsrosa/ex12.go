package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.ToUpper(s)

	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversedStr := string(chars)

	if s == reversedStr {
		fmt.Print("A string é um palíndromo ")
	} else {
		fmt.Print("A string não é um palíndromo ")
	}
}
