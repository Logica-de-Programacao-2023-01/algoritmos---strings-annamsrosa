package main

import (
	"fmt"
	"strings"
)

func main() {

	var str, char, troca string
	fmt.Print("Digite algo: ")
	fmt.Scan(&str)

troca:
	fmt.Print("Digite o caractere que será substituido: ")
	fmt.Scan(&troca)
	if strings.Contains(str, troca) == true {
		fmt.Print("Digite o caractere que substituirá o retirado: ")
		fmt.Scan(&char)
	} else {
		fmt.Print("O caractere informado não pertence à string!\n")
		goto troca
	}

	NS := strings.ReplaceAll(str, troca, char)

	fmt.Print("A nova string fica: ", NS)
}
