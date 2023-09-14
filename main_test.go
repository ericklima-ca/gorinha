package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type testFiles struct {
	fname string
	path  string
}

const DIR = "./files/"

var tests = []testFiles{
	{"print", fmt.Sprintf("%s%s", DIR, "print.json")},
	{"fib", fmt.Sprintf("%s%s", DIR, "fib.json")},
	{"sum", fmt.Sprintf("%s%s", DIR, "sum.json")},
	{"sub", fmt.Sprintf("%s%s", DIR, "sub.json")},
	{"concate", fmt.Sprintf("%s%s", DIR, "concate.json")},
}

func _eval(path string) interface{} {
	f, _ := os.ReadFile(path)
	var file File
	json.Unmarshal(f, &file)
	var scope = make(map[string]interface{})

	r := eval(file.Expression, scope)
	return r
}

func Test(t *testing.T) {
	for _, v := range tests {
		if r := _eval(v.path); r != nil {
			t.Errorf("error on %s.json file", v.fname)
		}
	}
}
