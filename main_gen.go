package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// EncodeMsg implements msgp.Encodable
func (z *ColoredPoint) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(3)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.X)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Y)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Color)
	if err != nil {
		return
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ColoredPoint) DecodeMsg(dc *msgp.Reader) (err error) {
	var ssz uint32
	ssz, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if ssz != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: ssz}
		return
	}
	z.X, err = dc.ReadFloat64()
	if err != nil {
		return
	}
	z.Y, err = dc.ReadFloat64()
	if err != nil {
		return
	}
	err = dc.ReadExtension(&z.Color)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ColoredPoint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, 3)
	o = msgp.AppendFloat64(o, z.X)
	o = msgp.AppendFloat64(o, z.Y)
	o, err = msgp.AppendExtension(o, &z.Color)
	if err != nil {
		return
	}
	return
}

//UnmarshalMsg implements msgp.Unmarshaler
func (z *ColoredPoint) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	}
	z.X, bts, err = msgp.ReadFloat64Bytes(bts)
	if err != nil {
		return
	}
	z.Y, bts, err = msgp.ReadFloat64Bytes(bts)
	if err != nil {
		return
	}
	bts, err = msgp.ReadExtensionBytes(bts, &z.Color)
	if err != nil {
		return
	}
	o = bts
	return
}

func (z *ColoredPoint) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + msgp.Float64Size + msgp.Float64Size + msgp.ExtensionPrefixSize + z.Color.Len()
	return
}
