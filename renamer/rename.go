package renamer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// BulkRenamer - função para renomear diversos arquivos de uma vez
func BulkRenamer(caminho, novoNome string) (bool, error) {
	contador := 0
	files, e := ioutil.ReadDir(caminho)

	showAllFiles(files)

	fmt.Println("Deseja continuar com a renomeação ?(s, sim ou deixe em branco para continuar)")
	reader := bufio.NewReader(os.Stdin)
	continuar, _ := reader.ReadString('\n')

	continuar = strings.ToLower(strings.TrimSpace(continuar))

	if continuar != "s" && continuar != "sim" && continuar != "" {
		return true, e
	}

	//TODO: Opção de selecionar extensões específicas para serem renomeadas, em branco, renomeará tudo.
	//NOTE: Usuário deve digitar as extensoes q quer utilizar e separar por virgula
	if e != nil {
		return false, e
	}

	for _, file := range files {
		var arquivoModificado string

		ext := filepath.Ext(file.Name())
		caminhoArquivo := fmt.Sprintf("%s/%s", caminho, file.Name())

		if novoNome == "" {
			arquivoModificado = fmt.Sprintf("%s/%04d%s", caminho, contador, ext)
		} else {
			arquivoModificado = fmt.Sprintf("%s/%d-%s%s", caminho, contador, novoNome, ext)
		}

		e := os.Rename(caminhoArquivo, arquivoModificado)

		if e != nil {
			return false, e
		}
		contador++
	}

	startExplorerOpt(caminho)

	return true, nil
}

// SingleRenamer - renomeia um único arquivo
func SingleRenamer(caminho, novoNome string) (bool, error) {
	base := filepath.Dir(caminho)
	ext := filepath.Ext(caminho)
	novoArquivo := fmt.Sprintf("%s/%s%s", base, novoNome, ext)

	e := os.Rename(caminho, novoArquivo)

	if e != nil {
		return false, e
	}

	startExplorerOpt(base)

	return true, nil
}

func startExplorerOpt(caminho string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Deseja abrir o explorer ? (y/n)")
	opt, _ := reader.ReadString('\n')

	opt = strings.TrimSpace(strings.ToLower(opt))

	if opt == "y" {
		openExplorer(caminho)
	}
}

func openExplorer(caminho string) {
	if runtime.GOOS == "linux" {
		exec.Command("xdg-open", caminho).Run()
	} else if runtime.GOOS == "windows" {
		exec.Command("explorer", caminho).Run()
	} else if runtime.GOOS == "darwin" {
		exec.Command("open", caminho).Run()
	} else {
		fmt.Println("Não foi possível iniciar o file explorer")
	}
}

func showAllFiles(files []os.FileInfo) {
	var lista []string

	for _, file := range files {
		lista = append(lista, file.Name())
	}

	fmt.Printf("Os seguinte arquivos serão renomeados: (%d arquivos)\n", len(lista))

	for index, arch := range lista {
		fmt.Printf("%d - %s\n", index, arch)
	}

	fmt.Println()
}
