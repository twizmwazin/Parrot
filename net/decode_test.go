package net

import (
	"bytes"
	p "github.com/twizmwazin/Parrot/proto"
	ty "github.com/twizmwazin/Parrot/types"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	var a1 uint8
	b1 := bytes.NewBuffer([]byte{0xc7}) // 199
	i1, err := DecodeType(a1, b1)
	a1 = i1.(uint8)
	if err != nil || a1 != 199 {
		t.FailNow()
	}

	var a2 uint16
	b2 := bytes.NewBuffer([]byte{0xde, 0xad}) // 57005
	i2, err := DecodeType(a2, b2)
	a2 = i2.(uint16)
	if err != nil || a2 != 57005 {
		t.FailNow()
	}

	var a9 string
	s := "Hello World"
	l := len(s)
	b9 := bytes.NewBuffer([]byte{})
	vl := ty.VarInt(l)
	vp := &vl

	err = vp.Write(b9)
	if err != nil {
		t.FailNow()
	}
	b9.WriteString(s)

	i9, err := DecodeType(a9, b9)
	a9 = i9.(string)
	if err != nil || strings.Compare(a9, "Hello World") != 0 {
		t.Errorf("'%s'", a9)
		t.FailNow()
	}
}

func TestDecodePayload(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	protoVer := ty.VarInt(123)
	err := protoVer.Write(buf)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	serverAddr := "test.com"
	serverAddrLen := ty.VarInt(len(serverAddr))
	serverAddrLen.Write(buf)
	buf.WriteString(serverAddr)

	b, err := EncodeType(uint16(25565))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	buf.Write(b)

	state := ty.VarInt(1)
	state.Write(buf)

	hs := p.HandshakeServerHandshake{}
	err = DecodePayload(&hs, buf)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if hs.ProtoVersion != 123 {
		t.Errorf("Invalid proto string, got %d, should be 123", hs.ProtoVersion)
		t.FailNow()
	}
	if strings.Compare(hs.ServerAddr, serverAddr) != 0 {
		t.Errorf("Invalid server address")
		t.FailNow()
	}
}
