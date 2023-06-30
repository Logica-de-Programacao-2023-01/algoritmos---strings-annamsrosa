package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite algo: ")
	fmt.Scanln(&str)

	inversão := ""
	for i := len(str) - 1; i >= 0; i-- {
		inversão += string(str[i])
	}

	fmt.Println("A inversão fica: ", inversão)
}
