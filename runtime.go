package main

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Runtime struct{}
type RuntimeError struct{}

func (r *Runtime) eval(e Term, scope Scope) Term {
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
			lv := reflect.TypeOf(r.eval(b.Lhs, scope)).Kind()
			rv := reflect.TypeOf(r.eval(b.Rhs, scope)).Kind()
			if lv == reflect.Int32 && rv == reflect.Int32 {
				return r.eval(b.Lhs, scope).(int32) + r.eval(b.Rhs, scope).(int32)
			} else if lv == reflect.String || rv == reflect.String {
				return fmt.Sprintf("%v%v", r.eval(b.Lhs, scope), r.eval(b.Rhs, scope))
			} else {
				panic("error add")
			}

		case "Sub":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv - rv
			} else {
				panic("error sub")
			}

		case "Mul":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv * rv
			} else {
				panic("error mul")
			}

		case "Div":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv / rv
			} else {
				panic("error div")
			}

		case "Rem":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv % rv
			} else {
				panic("error rem")
			}

		case "Lt":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv < rv
			} else {
				panic("error lt")
			}

		case "Lte":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv <= rv
			} else {
				panic("error lte")
			}

		case "Gt":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv > rv
			} else {
				panic("error gt")
			}

		case "Gte":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv >= rv
			} else {
				panic("error gte")
			}

		case "Eq":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv == rv
			} else {
				panic("error eq")
			}

		case "Neq":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv != rv
			} else {
				panic("error neq")
			}

		case "Or":
			lv, okl := r.eval(b.Lhs, scope).(bool)
			rv, okr := r.eval(b.Rhs, scope).(bool)
			if okl && okr {
				return lv || rv
			} else {
				panic("error or")
			}

		case "And":
			lv, okl := r.eval(b.Lhs, scope).(bool)
			rv, okr := r.eval(b.Rhs, scope).(bool)
			if okl && okr {
				return lv && rv
			} else {
				panic("error and")
			}

		}
	case "If":
		var i If
		decode(e, &i)
		if r.eval(i.Condition, scope).(bool) {
			return r.eval(i.Then, scope)
		} else {
			return r.eval(i.Otherwise, scope)
		}

	case "Let":
		var l Let
		decode(e, &l)
		scope[l.Name.Text] = r.eval(l.Value, scope)
		return r.eval(l.Next, scope)

	case "Var":
		var v Var
		decode(e, &v)
		return scope[v.Text]

	case "Function":
		var f Function
		decode(e, &f)
		return func(args []Term, fScope map[string]interface{}) Term {
			if len(args) != len(f.Parameters) {
				return RuntimeError{}
			}
			h := map[string]interface{}{}
			for k, v := range fScope {
				h[k] = v
			}
			for i, v := range f.Parameters {
				h[v.Text] = args[i]
			}
			return r.eval(f.Value, h)
		}

	case "Call":
		var c Call
		decode(e, &c)
		f := r.eval(c.Callee, scope)
		fn := reflect.ValueOf(f)

		var evalArgs []Term
		for _, v := range c.Arguments {
			evalArgs = append(evalArgs, r.eval(v, scope))
		}
		return fn.Call([]reflect.Value{reflect.ValueOf(evalArgs), reflect.ValueOf(scope)})[0].Interface().(Term)

	case "Print":
		var p Print
		decode(e, &p)
		v := r.eval(p.Value, scope)
		switch t := v.(type) {
		case Tuple:
			fmt.Printf("(%v, %v)\n", r.eval(t.First, scope), r.eval(t.Second, scope))
		default:
			if reflect.ValueOf(v).Kind() == reflect.Func {
				fmt.Println("<#closure>")
			} else {
				fmt.Println(v)
			}
		}
		return v

	case "Tuple":
		var t Tuple
		decode(e, &t)
		return t

	case "First":
		var f First
		decode(e, &f)
		v := r.eval(f.Value, scope).(Tuple).First
		return r.eval(v, scope)

	case "Second":
		var s Second
		decode(e, &s)
		v := r.eval(s.Value, scope).(Tuple).Second
		return r.eval(v, scope)
	}

	return RuntimeError{}
}

func decode(i interface{}, o interface{}) {
	mapstructure.Decode(i, &o)
}
