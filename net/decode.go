package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/MJKWoolnough/byteio"
	t "github.com/twizmwazin/Parrot/types"
	"io"
	"reflect"
)

// DecodePacket reads the next packet from r and returns the id, payload, or
// error, if applicable. compress specifies if the protocol is encrypted.
func DecodePacket(r io.ByteReader, pmap *map[int32]interface{}, compress bool) (id int32, payload interface{}, err error) {
	// Packet length
	var plen t.VarInt
	plen, err = t.NewVarInt(r)
	if err != nil {
		return
	}

	// Grab rest of packet before decoding
	rem := make([]byte, int(plen))
	for i := 0; i < int(plen); i++ {
		rem[i], err = r.ReadByte()
		if err != nil {
			return
		}
	}
	rbuf := bytes.NewBuffer(rem)

	var pid t.VarInt
	if compress {
		// TODO
		panic("compressed decoder not yet implemented")
	} else {
		pid, err = t.NewVarInt(rbuf)
		if err != nil {
			return
		}
		id = int32(pid)
		payload = rbuf.Bytes()
		// TODO: get struct from
		payload = (*pmap)[id]
		err = DecodePayload(&payload, rbuf)
	}

	return
}

// DecodePayload decodes the payload of a packet of type pt, into pt.
// Returns an error if the packet is unable to be decoded, otherwise nil.
func DecodePayload(p interface{}, buf *bytes.Buffer) (err error) {
	var in interface{}
	pt := reflect.ValueOf(p)
	t := pt.Elem()

	for i := 0; i < t.NumField(); i++ {
		vi := t.Field(i).Interface()
		// fmt.Println("vi is of type:", vi.Type(), vi)
		in, err = DecodeType(vi, buf)
		if err != nil {
			return
		}
		fmt.Println(vi)
		t.Field(i).Set(reflect.ValueOf(in))
	}

	return
}

// DecodeType decodes a value from buf of the type of t. Returns the value and
// an error if the type was unable to be decoded.
func DecodeType(pt interface{}, buf *bytes.Buffer) (r interface{}, err error) {
	be := byteio.BigEndianReader{Reader: buf}
	switch c := pt.(type) {
	case int8:
		r, _, err = be.ReadInt8()
	case int16:
		r, _, err = be.ReadInt16()
	case int32:
		r, _, err = be.ReadInt32()
	case int64:
		r, _, err = be.ReadInt64()
	case uint8:
		r, _, err = be.ReadUint8()
	case uint16:
		r, _, err = be.ReadUint16()
	case uint32:
		r, _, err = be.ReadUint32()
	case uint64:
		r, _, err = be.ReadUint64()
	case string:
		var l t.VarInt
		l, err = t.NewVarInt(buf)
		if err != nil {
			return
		}
		raw := make([]byte, int(l))
		for i := 0; i < len(raw); i++ {
			raw[i], err = buf.ReadByte()
			if err != nil {
				return
			}
		}
		r = string(raw)
	case t.VarInt:
		err = c.Read(buf)
		r = c
	case t.VarLong:
		err = c.Read(buf)
		r = c
	default:
		err = errors.New("unsupported type")
		fmt.Println("unsupported type", reflect.TypeOf(pt))
	}

	return
}
