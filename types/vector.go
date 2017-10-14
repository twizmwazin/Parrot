package types

// Vector3i represents an immutable 3-dimensional vector.
type Vector3i struct {
	x int64
	y int64
	z int64
}

// X returns the x-coordinate of the vector.
func (v Vector3i) X() int64 {
	return v.x
}

// Y returns the y-coordinate of the vector.
func (v Vector3i) Y() int64 {
	return v.y
}

// Z returns the z-coordinate of the vector.
func (v Vector3i) Z() int64 {
	return v.z
}

// Negate multiplies all values by -1.
func (v Vector3i) Negate() Vector3i {
	return v.Mul(-1)
}

// Add adds the specified vector and the called vector.
func (v Vector3i) Add(o Vector3i) Vector3i {
	return Vector3i{v.x + o.x, v.y + o.y, v.z + o.z}
}

// Sub subtracts the specified vector from the called vector.
func (v Vector3i) Sub(o Vector3i) Vector3i {
	return v.Add(o.Negate())
}

// Mul multiplies the called vector by the scalar value i.
func (v Vector3i) Mul(i int64) Vector3i {
	return Vector3i{v.x * i, v.y * i, v.z * i}
}

// Div divides the called vector by the scalar value i.
// The function assumes i does not contain any zero values.
func (v Vector3i) Div(i int64) Vector3i {
	return Vector3i{v.x / i, v.y / i, v.z / i}
}
