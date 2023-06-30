package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Print("Digite algo: ")
	fmt.Scanln(&s)

	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversãoS := string(chars)

	fmt.Print("A inversão ficou assim: ", reversãoS)
}
