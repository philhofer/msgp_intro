package main

import (
	"fmt"
)

//go:generate msgp

//msgp:tuple Thing

//msgp:shim Enum as:string using:(Enum).String/efromStr

type Enum int8

const (
	A Enum = iota
	B
	invalid Enum = -1
)

func (e Enum) String() string {
	switch e {
	case A:
		return "A"
	case B:
		return "B"
	default:
		return fmt.Sprintf("Enum(%d)", e)
	}
}

func efromStr(s string) Enum {
	switch s {
	case "A":
		return A
	case "B":
		return B
	default:
		return invalid
	}
}

type Thing struct {
	E Enum `msg:"enum"`
}
