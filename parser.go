package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Parser struct{}

func (p *Parser) parse() Term {
	args := os.Args
	if len(args) != 2 {
		panic(fmt.Errorf("command unknown"))
	}
	f, err := os.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	var file File

	json.Unmarshal(f, &file)

	return file.Expression
}
