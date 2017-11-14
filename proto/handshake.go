package proto

import (
	t "github.com/twizmwazin/Parrot/types"
)

type HandshakeServerHandshake struct {
	ProtoVersion t.VarInt
	ServerAddr   string
	ServerPort   uint16
	NextState    t.VarInt
}

type HandshakeServerHandshakeLegacy struct {
	Payload uint8
}
