package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scan(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scan(&str2)
	if str1 == str2 {
		fmt.Print("As strings são iguais ")
	} else {
		fmt.Print("As strings são diferentes ")
	}
}
