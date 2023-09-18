package main

type BinaryOp string

type Scope map[string]interface{}

type Term interface{}

type Binary struct {
	Lhs      Term     `json:"lhs"`
	Rhs      Term     `json:"rhs"`
	Op       BinaryOp `json:"op"`
	Location Location `json:"location"`
}

type If struct {
	Condition Term     `json:"condition"`
	Then      Term     `json:"then"`
	Otherwise Term     `json:"otherwise"`
	Location  Location `json:"location"`
}

type Call struct {
	Callee    Term     `json:"callee"`
	Arguments []Term   `json:"arguments"`
	Location  Location `json:"location"`
}

type Var struct {
	Text     string   `json:"text"`
	Location Location `json:"location"`
}
type Bool struct {
	Value    bool     `json:"value"`
	Location Location `json:"location"`
}

type Int struct {
	Value    int32    `json:"value"`
	Location Location `json:"location"`
}

type Tuple struct {
	First    Term     `json:"first"`
	Second   Term     `json:"second"`
	Location Location `json:"location"`
}

type Str struct {
	Value    string   `json:"value"`
	Location Location `json:"location"`
}

type Let struct {
	Name     Parameter `json:"name"`
	Next     Term      `json:"next"`
	Value    Term      `json:"value"`
	Location Location  `json:"location"`
}

type Function struct {
	Value      Term        `json:"value"`
	Parameters []Parameter `json:"parameters"`
	Location   Location    `json:"location"`
}

type Parameter struct {
	Text     string   `json:"text"`
	Location Location `json:"location"`
}

type Print struct {
	Value    Term     `json:"value"`
	Location Location `json:"location"`
}

type First struct {
	Value    Term     `json:"value"`
	Location Location `json:"location"`
}

type Second struct {
	Value    Term     `json:"value"`
	Location Location `json:"location"`
}

type Location struct {
	Start    uint32 `json:"start"`
	End      uint32 `json:"end"`
	Filename string `json:"filename"`
}

type File struct {
	Name       string   `json:"name"`
	Expression Term     `json:"expression"`
	Location   Location `json:"location"`
}
