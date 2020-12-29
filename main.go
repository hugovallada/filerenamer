package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hugovallada/filerenamer/renamer"
)

func main() {
	var caminho string
	var novoNome string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bem vindo ao FileRenamer - Versão atual: 2.0")

	fmt.Println("Deseja trabalhar com um arquivo(f) ou com um diretório(d/dir)?")
	opt, _ := reader.ReadString('\n')

	opt = strings.ToLower(strings.TrimSpace(opt))

	if opt == "d" || opt == "dir" {
		fmt.Println("Qual o diretório onde os arquivos serão renomeados ?")
		caminho, _ = reader.ReadString('\n')
		
		

		fmt.Println("Qual o nome base do arquivos ? Deixe em branco para utilizar um contador")
		novoNome, _ = reader.ReadString('\n')

		caminho, _ = filepath.Abs(strings.TrimSpace(renamer.LinuxHomeDirectoryReplace(caminho)))
		novoNome = strings.TrimSpace(novoNome)

		renamer.BulkRenamer(caminho, novoNome)

	} else if opt == "f" {
		fmt.Println("Qual o caminho do arquivo a ser renomeado ?")
		caminho, _ = reader.ReadString('\n')

		fmt.Println("Qual o novo nome?")
		novoNome, _ = reader.ReadString('\n')

		caminho, _ = filepath.Abs(strings.TrimSpace(caminho))
		novoNome = strings.TrimSpace(novoNome)

		renamer.SingleRenamer(caminho, novoNome)
	} else {
		fmt.Println("Você não selecionou uma opção válida")
		main()
	}
}
