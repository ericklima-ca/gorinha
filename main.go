package main

import (
	"log"
)

func main() {
	parser := Parser{}
	expression := parser.parse()
	runtime := Runtime{}

	var scope = make(Scope)

	_, err := runtime.eval(expression, scope).(RuntimeError)
	if err {
		log.Fatalln("\n===\nError occurred in main thread\n===")
	}
}
