# filerenamer

## Ferramenta para renomeação em lote de arquivos em um mesmo diretório, ou de um arquivo único

## Como funciona:


    Opção: Indicar se é irá trabalhar em lote/diretório (dir) ou com arquivo(file)

    DIR:

    Caminho: Indicar o caminho do diretório

    Nome Base: Indicar qual o nome base para os seus arquivos, caso fique em branco, utilizará apenas o contador de arquivos como nome
    O formato do nome será XXXX-nomeBase.ext ou XXXX.ext

    Extensões: Indicar as extensões dos arquivos que devem ser renomeados, caso nenhuma seja passada, todas as extensões no diretório serão usadas para renomeação

    Continuar: Após a listagem de todas as extensões que serão renomeadas, o usuário pode cancelar ou continuar a renomeação em lote

    FILE:

    Caminho: Indicar o caminho do arquivo

    Nome: Indicar o novo nome que o arquivo terá

## Como executar:

    1 º Baixar o projeto via zip ou com : git clone https://github.com/hugovallada/filerenamer.git;
    2 º Caso necessário, extraia a pasta;
    3 º Acessar a pasta do projeto via terminal;
    4 º De permissão de execussão ao arquivo binário filerenamer, se necessário;
    5 º Execute no terminal o comando ./filerenamer
