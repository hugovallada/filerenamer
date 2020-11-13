package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Deseja trabalhar com um arquivo(f) ou com um diretório(d/dir)?")
	opt, _ := reader.ReadString('\n')

	opt = strings.ToLower(strings.TrimSpace(opt))

	if opt == "d" || opt == "dir" {
		fmt.Println("Você selecionou a opção de trabalhar com diretório")
	} else if opt == "f" {
		fmt.Println("Você selecionou trabalhar com arquivos")
	} else {
		fmt.Println("Você não selecionou uma opção válida")
	}

}
