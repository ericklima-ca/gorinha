# Gorinha :thumbsup:
Interpreter for the [Rinha](https://github.com/aripiprazole/rinha-de-compiler) language's Abstract Syntax Tree (AST).  
Developed hastily for the challenge 😢.

### Build
```
go build -o ./bin/gorinha .as long as "rinha" is installed and accessible in $PATH.
```
Builds the program and saves it in the [bin](./bin/) directory.

The `gorinha` binary accepts the AST in `json` format or a `.rinha` file, provided that `rinha` is installed and accessible in the `$PATH`.

### Testing
The initial test files are located in the [files](./files/) directory.
To run tests:
```sh
./bin/gorinha files/print.json
# or
./bin/gorinha files/fib.rinha
```
or
```sh
docker build -t gorinha .
docker run gorinha files/print.json
```
