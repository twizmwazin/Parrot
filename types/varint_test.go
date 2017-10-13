package types

import (
	"bytes"
	"fmt"
	"testing"
)

var intCases map[int32][]byte
var longCases map[int64][]byte

func setupInt() {
	intCases = make(map[int32][]byte)
	intCases[0] = []byte{0x00}
	intCases[1] = []byte{0x01}
	intCases[2] = []byte{0x02}
	intCases[127] = []byte{0x7f}
	intCases[128] = []byte{0x80, 0x01}
	intCases[255] = []byte{0xff, 0x01}
	intCases[2147483647] = []byte{0xff, 0xff, 0xff, 0xff, 0x07}
	intCases[-1] = []byte{0xff, 0xff, 0xff, 0xff, 0x0f}
	intCases[-2147483648] = []byte{0x80, 0x80, 0x80, 0x80, 0x08}
}

func setupLong() {
	longCases = make(map[int64][]byte)
	longCases[0] = []byte{0x00}
	longCases[1] = []byte{0x01}
	longCases[2] = []byte{0x02}
	longCases[127] = []byte{0x7f}
	longCases[128] = []byte{0x80, 0x01}
	longCases[255] = []byte{0xff, 0x01}
	longCases[2147483647] = []byte{0xff, 0xff, 0xff, 0xff, 0x07}
	longCases[9223372036854775807] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}
	longCases[-1] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	longCases[-2147483648] = []byte{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01}
	longCases[-9223372036854775808] = []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}
}

func testVIRead(t *testing.T, b []byte, e int32) {
	n, err := NewVarInt(bytes.NewBuffer(b))
	if err != nil || n.Value() != e {
		t.FailNow()
	}
}

func testVIWrite(t *testing.T, in int32, e []byte) {
	var buf = new(bytes.Buffer)
	v := VarInt(in)
	v.Write(buf)
	if buf.Len() != len(e) {
		fmt.Println(in)
		fmt.Println(buf.Len())
		fmt.Println(len(e))
		t.FailNow()
	}
	for i := 0; i < len(e); i++ {
		if buf.Bytes()[i] != e[i] {
			t.FailNow()
		}
	}
}

func testVLRead(t *testing.T, b []byte, e int64) {
	n, err := NewVarLong(bytes.NewBuffer(b))
	if err != nil || n.Value() != e {
		t.FailNow()
	}
}

func testVLWrite(t *testing.T, in int64, e []byte) {
	var buf = new(bytes.Buffer)
	v := VarLong(in)
	v.Write(buf)
	if buf.Len() != len(e) {
		fmt.Println(in)
		fmt.Println(buf.Len())
		fmt.Println(len(e))
		t.FailNow()
	}
	for i := 0; i < len(e); i++ {
		if buf.Bytes()[i] != e[i] {
			t.FailNow()
		}
	}
}

func TestVarIntRead(t *testing.T) {
	setupInt()
	for k, v := range intCases {
		testVIRead(t, v, k)
	}
}

func TestVarIntWrite(t *testing.T) {
	setupInt()
	for k, v := range intCases {
		testVIWrite(t, k, v)
	}
}

func TestVarLongRead(t *testing.T) {
	setupLong()
	for k, v := range longCases {
		testVLRead(t, v, k)
	}
}

func TestVarLongWrite(t *testing.T) {
	setupLong()
	for k, v := range longCases {
		testVLWrite(t, k, v)
	}
}
