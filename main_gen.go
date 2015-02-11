package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// EncodeMsg implements msgp.Encodable
func (z *Blob) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("name")
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	err = en.WriteString("data")
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		return
	}
	err = en.WriteString("num")
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Num)
	if err != nil {
		return
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Blob) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				return
			}
		case "num":
			z.Num, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

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

// START GEN OMIT

// MarshalMsg implements msgp.Marshaler
func (z *Blob) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "name")
	o = msgp.AppendString(o, z.Name)
	o = msgp.AppendString(o, "data")
	o = msgp.AppendBytes(o, z.Data)
	o = msgp.AppendString(o, "num")
	o = msgp.AppendFloat64(o, z.Num)
	return
}

// END GEN OMIT

//UnmarshalMsg implements msgp.Unmarshaler
func (z *Blob) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				return
			}
		case "num":
			z.Num, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Blob) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.Name) + msgp.StringPrefixSize + 4 + msgp.BytesPrefixSize + len(z.Data) + msgp.StringPrefixSize + 3 + msgp.Float64Size
	return
}

// START CPTGEN OMIT

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

// END CPTGEN OMIT

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
