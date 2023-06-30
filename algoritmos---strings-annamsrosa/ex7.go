package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string
	var count int

	count = 0

	fmt.Print("Digite uma string:")
	fmt.Scanln(&str)

	for _, n := range str {
		if _, err := strconv.Atoi(string(n)); err == nil {
			count++
		}
	}
	if count > 0 {
		fmt.Print("A string possui um número ")
	} else {
		fmt.Print("A string não possui um número ")
	}
}
