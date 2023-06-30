package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var frase string

	fmt.Print("Digite uma frase: ")
	reader := bufio.NewReader(os.Stdin)
	frase, _ = reader.ReadString('\n')

	espaço := strings.ReplaceAll(frase, " ", "")
	fmt.Println("A frase fica: ", espaço)
}
