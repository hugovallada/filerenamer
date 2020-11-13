package renamer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	return true, nil
}
