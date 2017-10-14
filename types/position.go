package types

// Position is a type that represents a world position. It is encoded as a 64-bit
// integer in the protocol.
type Position Vector3i

// DecodePosition will take in an int64 and output a Position.
func DecodePosition(i uint64) Position {
	var x = i >> 38
	var y = (i >> 26) & 0xfff
	var z = i << 38 >> 38
	return Position{int64(x), int64(y), int64(z)}

}

// Encode will return a uint64 encoding of the position.
func (p Position) Encode() uint64 {
	return uint64(((p.x & 0x3ffffff) << 38) | ((p.y & 0xfff) << 26) | (p.z & 0x3ffffff))
}
