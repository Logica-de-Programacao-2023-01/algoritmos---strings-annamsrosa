package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	únicas := ""
	for _, char := range s {
		count := strings.Count(s, string(char))
		if count == 1 {
			únicas += string(char)
		}
	}
	fmt.Print("As letras únicas presentes na string são: ", únicas)
}
