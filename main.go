package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type BinaryOp string

type Term interface{}

type Tuple struct {
	First  Term `json:"first"`
	Second Term `json:"second"`
}

type Binary struct {
	Lhs Term     `json:"lhs"`
	Rhs Term     `json:"rhs"`
	Op  BinaryOp `json:"op"`
}

type If struct {
	Condition Term `json:"condition"`
	Then      Term `json:"then"`
	Otherwise Term `json:"otherwise"`
}

type Call struct {
	Callee    Term   `json:"callee"`
	Arguments []Term `json:"arguments"`
}

type Var struct {
	Text string `json:"text"`
}
type Bool struct {
	Value bool `json:"value"`
}

type Int struct {
	Value int32 `json:"value"`
}
type Str struct {
	Value string `json:"value"`
}

type Let struct {
	Name  Parameter `json:"name"`
	Next  Term      `json:"next"`
	Value Term      `json:"value"`
}

type Function struct {
	Value      Term        `json:"value"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Text string `json:"text"`
}

type Print struct {
	Value Term `json:"value"`
}

type Location struct {
	Start    uint32 `json:"start"`
	End      uint32 `json:"end"`
	Filename string `json:"filename"`
}
type File struct {
	Name       string `json:"name"`
	Expression Term   `json:"expression"`
}

var loop int32 = 0

func main() {
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

	e := file.Expression

	var scope = make(map[string]interface{})
	eval(e, scope)
}

func eval(e Term, scope map[string]interface{}) interface{} {
	if e == nil {
		panic("not file")
	}
	switch e.(map[string]interface{})["kind"] {
	case "Str":
		var s Str
		decode(e, &s)
		return s.Value
	case "Int":
		var s Int
		decode(e, &s)
		return s.Value
	case "Bool":
		var s Bool
		decode(e, &s)
		return s.Value
	case "Binary":
		var b Binary
		decode(e, &b)
		switch b.Op {
		case "Add":
			l := reflect.TypeOf(eval(b.Lhs, scope)).Kind()
			r := reflect.TypeOf(eval(b.Rhs, scope)).Kind()
			if l == reflect.Int32 && r == reflect.Int32 {
				return eval(b.Lhs, scope).(int32) + eval(b.Rhs, scope).(int32)
			} else if l == reflect.String || r == reflect.String {
				return fmt.Sprintf("%v%v", eval(b.Lhs, scope), eval(b.Rhs, scope))
			} else {
				panic("error add")
			}

		case "Sub":
			lv, okl := eval(b.Lhs, scope).(int32)
			rv, okr := eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv - rv
			} else {
				panic("error sub")
			}
		case "Lt":
			lv, okl := eval(b.Lhs, scope).(int32)
			rv, okr := eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv < rv
			} else {
				panic("error lt")
			}
		case "Eq":
			lv, okl := eval(b.Lhs, scope).(int32)
			rv, okr := eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv == rv
			} else {
				panic("error eq")
			}

		case "Neq":
			lv, okl := eval(b.Lhs, scope).(int32)
			rv, okr := eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv != rv
			} else {
				panic("error neq")
			}

		}
	case "If":
		var i If
		decode(e, &i)
		if eval(i.Condition, scope).(bool) {
			return eval(i.Then, scope)
		} else {
			return eval(i.Otherwise, scope)
		}
	case "Let":
		var l Let
		decode(e, &l)
		scope[l.Name.Text] = eval(l.Value, scope)
		return eval(l.Next, scope)

	case "Var":
		var v Var
		decode(e, &v)
		return scope[v.Text]

	case "Function":
		var f Function
		decode(e, &f)
		return func(args []Term, fScope map[string]interface{}) Term {
			h := map[string]interface{}{}
			for k, v := range fScope {
				h[k] = v
			}
			for i, v := range f.Parameters {
				h[v.Text] = args[i]
			}
			return eval(f.Value, h)
		}

	case "Call":
		var c Call
		decode(e, &c)
		f := eval(c.Callee, scope)
		fn := reflect.ValueOf(f)

		var evalArgs []Term
		for _, v := range c.Arguments {
			evalArgs = append(evalArgs, eval(v, scope))
		}
		return fn.Call([]reflect.Value{reflect.ValueOf(evalArgs), reflect.ValueOf(scope)})[0].Interface().(Term)

	case "Print":
		var p Print
		decode(e, &p)
		fmt.Println(eval(p.Value, scope))
	}
	return nil
}

func decode(i interface{}, o interface{}) {
	mapstructure.Decode(i, &o)
}
