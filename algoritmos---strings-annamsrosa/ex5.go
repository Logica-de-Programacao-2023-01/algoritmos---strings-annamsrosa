package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	var count int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	for _, char := range num {
		if string(char) == "." {
			count += 1
		}
	}

	if count == 1 {
		if _, err := strconv.ParseFloat(num, 64); err == nil {
			fmt.Println("Número válido para ponto flutuante ")
		} else {
			fmt.Println("Número inválido para ponto flutuante ")
		}

	} else {
		fmt.Println("Número inválido para ponto flutuante ")
	}
}
