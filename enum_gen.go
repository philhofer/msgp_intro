package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// EncodeMsg implements msgp.Encodable
func (z *Thing) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(1)
	if err != nil {
		return
	}
	err = en.WriteString((Enum).String(z.E))
	if err != nil {
		return
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Thing) DecodeMsg(dc *msgp.Reader) (err error) {
	var ssz uint32
	ssz, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if ssz != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: ssz}
		return
	}
	{
		var tmp string
		tmp, err = dc.ReadString()
		z.E = efromStr(tmp)
	}
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Thing) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, 1)
	o = msgp.AppendString(o, (Enum).String(z.E))
	return
}

//UnmarshalMsg implements msgp.Unmarshaler
func (z *Thing) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var ssz uint32
		ssz, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if ssz != 1 {
			err = msgp.ArrayError{Wanted: 1, Got: ssz}
			return
		}
	}
	{
		var tmp string
		tmp, bts, err = msgp.ReadStringBytes(bts)
		z.E = efromStr(tmp)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

func (z *Thing) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + msgp.StringPrefixSize + len((Enum).String(z.E))
	return
}
