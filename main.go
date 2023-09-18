package main

func main() {
	parser := Parser{}
	runtime := Runtime{}
	expression := parser.parse()
	var scope = make(Scope)
	_, err := runtime.eval(expression, scope).(RuntimeError)
	if err {
		panic("runtime error")
	}
}
