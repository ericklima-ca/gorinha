# Gorinha :thumbsup:
Interpretador da AST da linguagem [Rinha](https://github.com/aripiprazole/rinha-de-compiler).  
Desenvolvido às pressas para o desafio 😢.
### Build
```
$ go build -o ./bin/gorinha .
```
faz o build do programa e salva em `bin`.

O binário `gorinha` aceita a AST em `json` e o executa.
### Teste
Os arquivos iniciais de teste estão na pasta [files](./files/).
Para testar:
```sh
$ ./bin/gorinha files/print.json
```
ou
```
docker build -t gorinha .
docker run gorinha files/print.json
```