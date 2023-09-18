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
	{"combination", fmt.Sprintf("%s%s", DIR, "combination.json")},
	{"first", fmt.Sprintf("%s%s", DIR, "first.json")},
	{"second", fmt.Sprintf("%s%s", DIR, "second.json")},
	{"print_tuple", fmt.Sprintf("%s%s", DIR, "print_tuple.json")},
	{"print_function", fmt.Sprintf("%s%s", DIR, "print_function.json")},
}

func _eval(path string) bool {
	f, _ := os.ReadFile(path)
	var file File
	json.Unmarshal(f, &file)
	var scope = make(map[string]interface{})

	runtime := Runtime{}

	_, err := runtime.eval(file.Expression, scope).(RuntimeError)
	return err
}

func Test(t *testing.T) {
	for _, v := range tests {
		if err := _eval(v.path); err {
			t.Errorf("error on %s.json file", v.fname)
		}
	}
}
