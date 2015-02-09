package main

import (
	"errors"
	"fmt"
	"github.com/tinylib/msgp/msgp"
	"os"
)

func init() {
	msgp.RegisterExtension(94, func() msgp.Extension { return &RGBA{} })
}

//go:generate msgp

//msgp:tuple ColoredPoint

// START OMIT
type RGBA [4]uint8

func (r RGBA) ExtensionType() int8 { return 94 }
func (r *RGBA) Len() int           { return 4 }

func (r *RGBA) MarshalBinaryTo(b []byte) error {
	if copy(b, (*r)[:]) != 4 {
		panic("misbehaving!")
	}
	return nil
}

func (r *RGBA) UnmarshalBinary(b []byte) error {
	if copy((*r)[:], b) != 4 {
		return errors.New("wrong size!")
	}
	return nil
}

// END OMIT

func (r *RGBA) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%d, %d, %d, %d]", r[0], r[1], r[2], r[3])), nil
}

type ColoredPoint struct {
	X, Y  float64
	Color RGBA `msg:",extension"`
}

func main() {
	c := ColoredPoint{
		X:     1.02081,
		Y:     -1.209821,
		Color: RGBA{255, 255, 255, 0},
	}

	data, _ := c.MarshalMsg(nil)
	msgp.UnmarshalAsJSON(os.Stdout, data)
	print("\n")
}
