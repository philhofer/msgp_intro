package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements the msgp.Decodable interface
func (z *Primitive) DecodeMsg(dc *msgp.Reader) (err error) {

	{
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != 3 {
			err = msgp.ArrayError{Wanted: 3, Got: ssz}
			return
		}

		z.One, err = dc.ReadInt()

		if err != nil {
			return
		}

		z.Two, err = dc.ReadUint()

		if err != nil {
			return
		}

		z.Three, err = dc.ReadFloat64()

		if err != nil {
			return
		}

	}

	return
}

// EncodeMsg implements the msgp.Encodable interface
func (z *Primitive) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteArrayHeader(3)
	if err != nil {
		return
	}

	err = en.WriteInt(z.One)

	if err != nil {
		return
	}

	err = en.WriteUint(z.Two)

	if err != nil {
		return
	}

	err = en.WriteFloat64(z.Three)

	if err != nil {
		return
	}

	return
}

// MarshalMsg implements the msgp.Marshaler interface
func (z *Primitive) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendArrayHeader(o, 3)

	o = msgp.AppendInt(o, z.One)

	o = msgp.AppendUint(o, z.Two)

	o = msgp.AppendFloat64(o, z.Three)

	return
}

// UnmarshalMsg unmarshals a Primitive from MessagePack, returning any extra bytes
// and any errors encountered
func (z *Primitive) UnmarshalMsg(bts []byte) (o []byte, err error) {

	{
		var ssz uint32
		ssz, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if ssz != 3 {
			err = msgp.ArrayError{Wanted: 3, Got: ssz}
			return
		}

		z.One, bts, err = msgp.ReadIntBytes(bts)

		if err != nil {
			return
		}

		z.Two, bts, err = msgp.ReadUintBytes(bts)

		if err != nil {
			return
		}

		z.Three, bts, err = msgp.ReadFloat64Bytes(bts)

		if err != nil {
			return
		}

	}

	o = bts
	return
}

// Msgsize implements the msgp.Sizer interface
func (z *Primitive) Msgsize() (s int) {

	s += msgp.ArrayHeaderSize

	s += msgp.IntSize

	s += msgp.UintSize

	s += msgp.Float64Size

	return
}
