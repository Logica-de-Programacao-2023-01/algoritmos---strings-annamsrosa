package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	maiúsculas := strings.ToUpper(palavra)

	fmt.Print("A palavra fica: ", maiúsculas)
}
