package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s string

	fmt.Print("Digite uma sequência numérica: ")
	fmt.Scanln(&s)

	decrescente := true
	for i := 1; i < len(s); i++ {
		nAnterior, err1 := strconv.Atoi(string(s[i-1]))
		nAtual, err2 := strconv.Atoi(string(s[i]))

		if err1 != nil || err2 != nil {
			fmt.Println("Sequência  inválida.")
			return
		}

		if nAnterior <= nAtual {
			decrescente = false
			break
		}
	}

	if decrescente {
		fmt.Println("A sequência é decrescente ")
	} else {
		fmt.Println("A sequência não é decrescente ")
	}
}
