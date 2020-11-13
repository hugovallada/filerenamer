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

	if e != nil {
		return false, e
	}

	for _, file := range files {
		var arquivoModificado string

		ext := filepath.Ext(file.Name())
		caminhoArquivo := fmt.Sprintf("%s/%s", caminho, file.Name())

		if novoNome == "" {
			arquivoModificado = fmt.Sprintf("%s/%d/%s", caminho, contador, ext)
		} else {
			arquivoModificado = fmt.Sprintf("%s/%d-%s%s", caminho, contador, novoNome, ext)
		}

		e := os.Rename(caminhoArquivo, arquivoModificado)

		if e != nil {
			return false, e
		}
		contador++
	}
	return true, nil
}
