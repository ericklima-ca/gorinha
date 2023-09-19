package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type Runtime struct {
	Filename string `json:"filename"`
}
type RuntimeError struct{}

func (r *Runtime) eval(e Term, scope Scope) Term {
	if e == nil {
		return RuntimeError{}
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
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv + rv
			}
			_, okls := left.(string)
			_, okrs := right.(string)
			if okls || okrs {
				return fmt.Sprintf("%v%v", left, right)
			} else {
				log.Fatalf("\n===\nError on add operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Sub":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv - rv
			} else {
				log.Fatalf("\n===\nError on sub operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Mul":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv * rv
			} else {
				log.Fatalf("\n===\nError on mul operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Div":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv / rv
			} else {
				log.Fatalf("\n===\nError on div operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Rem":
			lv, okl := r.eval(b.Lhs, scope).(int32)
			rv, okr := r.eval(b.Rhs, scope).(int32)
			if okl && okr {
				return lv % rv
			} else {
				log.Fatalf("\n===\nError on rem operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Lt":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv < rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs < rvs
			} else {
				log.Fatalf("\n===\nError on lt operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Lte":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv <= rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs <= rvs
			} else {
				log.Fatalf("\n===\nError on lte operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Gt":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv > rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs > rvs
			} else {
				log.Fatalf("\n===\nError on gt operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Gte":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv >= rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs >= rvs
			} else {
				log.Fatalf("\n===\nError on gte operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Eq":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv == rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs == rvs
			} else {
				log.Fatalf("\n===\nError on eq operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Neq":
			left := r.eval(b.Lhs, scope)
			right := r.eval(b.Rhs, scope)
			lv, okl := left.(int32)
			rv, okr := right.(int32)
			if okl && okr {
				return lv != rv
			}
			lvs, okls := left.(string)
			rvs, okrs := right.(string)
			if okls && okrs {
				return lvs != rvs
			} else {
				log.Fatalf("\n===\nError on neq operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "Or":
			lv, okl := r.eval(b.Lhs, scope).(bool)
			rv, okr := r.eval(b.Rhs, scope).(bool)
			if okl && okr {
				return lv || rv
			} else {
				log.Fatalf("\n===\nError on or operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
			}

		case "And":
			lv, okl := r.eval(b.Lhs, scope).(bool)
			rv, okr := r.eval(b.Rhs, scope).(bool)
			if okl && okr {
				return lv && rv
			} else {
				log.Fatalf("\n===\nError on and operator at %s\n===\n", strconv.Itoa(int(b.Location.Start)))
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
		return func(args []Term, fScope Scope) Term {
			if len(args) != len(f.Parameters) {
				return RuntimeError{}
			}
			for i, v := range f.Parameters {
				scope[v.Text] = args[i]
			}
			newScope := make(Scope)
			for k, v := range scope {
				newScope[k] = v
			}
			return r.eval(f.Value, newScope)
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
