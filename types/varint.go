package types

import (
	"encoding/binary"
	"io"
)

// VarInt represents a 32-bit varible length int
type VarInt int32

func (v *VarInt) Read(r io.ByteReader) (err error) {
	x, err := binary.ReadUvarint(r)
	if err != nil {
		return
	}
	*v = VarInt(x)
	return
}

// Adapted from standard library to work with 32-bit integers
func putUvarint32(buf []byte, x uint32) int {
	i := 0
	for x >= 0x80 {
		buf[i] = byte(x) | 0x80
		x >>= 7
		i++
	}
	buf[i] = byte(x)
	return i + 1
}

func (v *VarInt) Write(w io.ByteWriter) (err error) {
	buf := make([]byte, 5)
	putUvarint32(buf, uint32(*v))
	for i, v := range buf {
		if i != 0 && v == 0 {
			break
		}
		err = w.WriteByte(v)
		if err != nil {
			return
		}
	}
	return
}

// Value returns the internal value as an int32
func (v *VarInt) Value() int32 {
	return int32(*v)
}

// NewVarInt creates a new VarInt
func NewVarInt(r io.ByteReader) (VarInt, error) {
	var v VarInt
	err := v.Read(r)
	return v, err
}

// VarLong represents a 64-bit varible length int
type VarLong int64

func (v *VarLong) Read(r io.ByteReader) (err error) {
	x, err := binary.ReadUvarint(r)
	if err != nil {
		return
	}
	*v = VarLong(x)
	return
}

func (v *VarLong) Write(w io.ByteWriter) (err error) {
	buf := make([]byte, 10)
	binary.PutUvarint(buf, uint64(*v))
	for i, v := range buf {
		if i != 0 && v == 0 {
			break
		}
		err = w.WriteByte(v)
		if err != nil {
			return
		}
	}
	return
}

// Value returns the internal value as an int64
func (v *VarLong) Value() int64 {
	return int64(*v)
}

// NewVarLong creates a new VarLong
func NewVarLong(r io.ByteReader) (VarLong, error) {
	var v VarLong
	err := v.Read(r)
	return v, err
}
