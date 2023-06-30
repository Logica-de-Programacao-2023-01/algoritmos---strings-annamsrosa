package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.Replace(s, "A", "", -1)
	s = strings.Replace(s, "a", "", -1)
	s = strings.Replace(s, "E", "", -1)
	s = strings.Replace(s, "e", "", -1)
	s = strings.Replace(s, "I", "", -1)
	s = strings.Replace(s, "i", "", -1)
	s = strings.Replace(s, "O", "", -1)
	s = strings.Replace(s, "o", "", -1)
	s = strings.Replace(s, "U", "", -1)
	s = strings.Replace(s, "u", "", -1)

	fmt.Println("A string fica: ", s)
}
