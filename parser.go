package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path"
)

type Parser struct{}

func (p *Parser) parse() Term {
	args := os.Args

	var filename string

	if len(args) < 2 {
		filename = "/var/rinha/source.rinha.json"
	} else {
		filename = args[1]
	}

	var outFile []byte

	ext := path.Ext(filename)

	if ext == ".rinha" {
		cmd := exec.Command("rinha", filename)
		out, err := cmd.CombinedOutput()
		if err != nil {
			errMessage := string(out)
			if errMessage == "" {
				message := "\n===\n\033[1;31mrinha not installed.\033[0m\n\nPlease install it by running: \033[1;34mcargo install rinha\033[1;0m\nOr try running \033[1mgorinha\033[0m on a .json file\n==="
				log.Fatalln(message)
			} else {
				log.Fatalln(errMessage)
			}
		}
		outFile = out
	} else if ext == ".json" {
		f, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("\n===\n\033[1mError on reading file.\033[0m\n===", err)
		}
		outFile = f
	} else {
		log.Fatalln("\n===\n\033[1mFile extension just .json or .rinha\033[0m\n===")
	}

	var file File

	json.Unmarshal(outFile, &file)

	return file.Expression
}
