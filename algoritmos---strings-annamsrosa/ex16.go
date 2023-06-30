package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s1)
	fmt.Print("Digite outra string: ")
	fmt.Scanln(&s2)

	if strings.Contains(s1, s2) == true {
		fmt.Print("A segunda é uma substring da primeira ")
	} else {
		fmt.Print("A segunda não é uma substring da primeira ")
	}
}
