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

	//TODO: Fazer a listagem dos arquivos a serem renomeados
	//TODO: Opção de ignorar arquivos de determinadas extensões

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

	//TODO: Opção de abrir o explorador de arquivos (Ambas as funções)
	openExplorer(caminho)

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

	//TODO: Perguntar ao usuário se ele deseja abrir o explorador
	opt := startExplorerOpt()

	if opt {
		openExplorer(base)
	}

	return true, nil
}

func startExplorerOpt() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Deseja abrir o explorer ? (y/n)")
	opt, _ := reader.ReadString('\n')

	opt = strings.TrimSpace(strings.ToLower(opt))

	if opt == "y" {
		return true
	}
	return false
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
