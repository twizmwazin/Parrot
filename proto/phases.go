package proto

type Protocol byte

const (
	Handshake Protocol = iota
	Status
	Login
	Play
)
