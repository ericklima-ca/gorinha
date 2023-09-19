# Gorinha :thumbsup:
Interpretador da AST da linguagem [Rinha](https://github.com/aripiprazole/rinha-de-compiler).  
Desenvolvido às pressas para o desafio 😢.
### Build
```
go build -o ./bin/gorinha .
```
faz o build do programa e salva em [bin](./bin/).

O binário `gorinha` aceita a AST em `json` e ou um arquivo `.rinha`, desde que `rinha` esteja instalado e acessível em `$PATH`.
### Teste
Os arquivos iniciais de teste estão na pasta [files](./files/).
Para testar:
```sh
./bin/gorinha files/print.json
# ou
./bin/gorinha files/fib.rinha
```
ou
```sh
docker build -t gorinha .
docker run gorinha files/print.json
```
