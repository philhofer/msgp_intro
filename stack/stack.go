package main

import (
	"github.com/tinylib/msgp/msgp"
	"os"
)

//go:generate msgp

//msgp:tuple Primitive

type Primitive struct {
	One   int
	Two   uint
	Three float64
}

func main() {
	p := Primitive{1, 2, 3.0}
	out, _ := p.MarshalMsg(nil)
	msgp.UnmarshalAsJSON(os.Stdout, out)
	print("\n")
}
