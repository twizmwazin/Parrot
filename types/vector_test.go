package types

import (
	"testing"
)

var (
	v    = Vector3i{1, 2, 3}
	w    = Vector3i{7, 8, 9}
	x    = Vector3i{-4, -5, -6}
	y    = Vector3i{2, 4, 6}
	zero = Vector3i{0, 0, 0}
)

func TestCreate(t *testing.T) {
	if v.X() != 1 || v.Y() != 2 || v.Z() != 3 {
		t.Fail()
	}
}

func TestNegate(t *testing.T) {
	if zero.Negate() != zero {
		t.Fail()
	}

	if (v.Negate() != Vector3i{-1, -2, -3}) {
		t.Fail()
	}

}

func TestAdd(t *testing.T) {
	if (v.Add(w) != Vector3i{8, 10, 12}) {
		t.Fail()
	}

	if (v.Add(x) != Vector3i{-3, -3, -3}) {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	if (v.Sub(w) != Vector3i{-6, -6, -6}) {
		t.Fail()
	}

	if (v.Sub(x) != Vector3i{5, 7, 9}) {
		t.Log(v.Sub(x))
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	if v.Mul(1) != v {
		t.Fail()
	}

	if v.Mul(0) != zero {
		t.Fail()
	}

	if (v.Mul(-1) != Vector3i{-1, -2, -3}) {
		t.Fail()
	}

	if (v.Mul(2) != Vector3i{2, 4, 6}) {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	if v.Div(1) != v {
		t.Fail()
	}

	if v.Div(-1) != v.Negate() {
		t.Fail()
	}

	if y.Div(2) != v {
		t.Fail()
	}

}
