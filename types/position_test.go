package types

import (
	"testing"
)

func TestPosition(t *testing.T) {
	var i1 uint64 = 0x1234567890abcdef
	var i2 uint64 = 0xfedcba0987654321
	p1 := DecodePosition(i1)
	p2 := DecodePosition(i2)
	r1 := p1.Encode()
	r2 := p2.Encode()

	if r1 != i1 || r2 != i2 {
		t.Fail()
	}
}
