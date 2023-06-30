package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, novo, substituição string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

troca:
	fmt.Print("Digite o caractere que será substituido: ")
	fmt.Scan(&substituição)
	if strings.Contains(str, substituição) == true {
		fmt.Print("Digite o caractere que substituirá o retirado: ")
		fmt.Scan(&novo)
	} else {
		fmt.Print("O caractere informado não pertence à string!\n")
		goto troca
	}

	nova := strings.ReplaceAll(str, substituição, novo)

	fmt.Print("A nova string ficou assim:", nova)
}
