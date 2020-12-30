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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite as extensões dos arquivos que quer renomear.(Deixar em branco renomeará todos os arquivos) - Separar por virgula")
	extensoes, _ := reader.ReadString('\n')

	if e != nil {
		return false, e
	}

	listExt := splitExtensions(extensoes)
	filesWithExt := saveFilesWithExtension(files, listExt)

	fmt.Printf("Os seguintes arquivos serão renomeados: (%d arquivos)\n", len(filesWithExt))
	for index, fileToRename := range filesWithExt {
		fmt.Printf("%d --- %s\n", index, fileToRename)
	}

	fmt.Println()

	fmt.Println("Deseja continuar com a renomeação ?(s, sim ou deixe em branco para continuar)")

	continuar, _ := reader.ReadString('\n')

	continuar = strings.ToLower(strings.TrimSpace(continuar))

	if continuar != "s" && continuar != "sim" && continuar != "" {
		return true, e
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

		if existsInSlice(ext, listExt) {
			e := os.Rename(caminhoArquivo, arquivoModificado)
			if e != nil {
				return false, e
			}
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

	fmt.Printf("O diretório possui os seguintes arquivos: (%d arquivos)\n", len(lista))

	for index, arch := range lista {
		fmt.Printf("%d - %s\n", index, arch)
	}

	fmt.Println()
}

func saveFilesWithExtension(files []os.FileInfo, extensions []string) []string {
	var listaFilesExt []string

	if checkIfSliceIsEmpty(extensions) {
		for _, file := range files {
			listaFilesExt = append(listaFilesExt, file.Name())
		}
		return listaFilesExt
	}

	for _, file := range files {
		insert := false
		for _, ext := range extensions {
			if filepath.Ext(file.Name()) == ext {
				insert = true
			}
		}

		if insert {
			listaFilesExt = append(listaFilesExt, file.Name())
		}
	}

	return listaFilesExt
}

func splitExtensions(extensions string) []string {
	var listExt []string

	listExtensions := strings.Split(extensions, ",")

	for _, ext := range listExtensions {
		listExt = append(listExt, strings.TrimSpace(strings.ToLower(ext)))
	}
	return listExt
}

func existsInSlice(searchValue string, searchableSlice []string) bool {

	if checkIfSliceIsEmpty(searchableSlice) {
		return true
	}

	for _, value := range searchableSlice {
		if searchValue == value {
			return true
		}
	}
	return false
}

func checkIfSliceIsEmpty(sliceCheck []string) bool {
	empty := true
	for _, val := range sliceCheck {
		if val != "" {
			empty = false
		}
	}
	return empty
}

// LinuxHomeDirectoryReplace exported
func LinuxHomeDirectoryReplace(caminho string) string {
	if caminho[0] == 126 && runtime.GOOS == "linux" {

		homeDir, _ := os.UserHomeDir()
		newCaminho := fmt.Sprintf("%s%s", homeDir, caminho[1:])

		return newCaminho
	}

	return caminho
}
