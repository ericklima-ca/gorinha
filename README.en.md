# Gorinha :thumbsup:
Interpreter for the [Rinha](https://github.com/aripiprazole/rinha-de-compiler) language's Abstract Syntax Tree (AST).  
Developed hastily for the challenge 😢.

### Build
```
$ go build -o ./bin/gorinha .
```
Builds the program and saves it in the `bin` directory.

The `gorinha` binary accepts the AST in `json` format and executes it.

### Testing
The initial test files are located in the [files](./files/) directory.
To run tests:
```sh
$ ./bin/gorinha files/print.json
```
or
```
docker build -t gorinha .
docker run gorinha files/print.json
```