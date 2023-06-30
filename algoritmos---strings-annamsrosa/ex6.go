package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase:")
	scanner.Scan()

	str1 := scanner.Text()
	str2 := strings.Count(str1, " ") + 1
	fmt.Print("A quantidade de palavras na frase é: ", str2)
}
