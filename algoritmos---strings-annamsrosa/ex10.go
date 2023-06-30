package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s1, s2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&s1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&s2)

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	chars1 := strings.Split(s1, "")
	chars2 := strings.Split(s2, "")

	sort.Strings(chars1)
	sort.Strings(chars2)

	sortedStr1 := strings.Join(chars1, "")
	sortedStr2 := strings.Join(chars2, "")

	if sortedStr2 == sortedStr1 {
		fmt.Println("As strings são anagramas ")
	} else {
		fmt.Println("As strings não são anagramas ")
	}
}
