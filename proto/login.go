package proto

import (
	t "github.com/twizmwazin/Parrot/types"
)

type LoginClientDisconnect struct {
	Reason string
}

type LoginClientEncyptionRequest struct {
	ServerID       string
	PubKeyLen      t.VarInt
	PubKey         []byte
	VerifyTokenLen t.VarInt
	VerifyToken    []byte
}

type LoginClientSetCompression struct {
	Threshold t.VarInt
}

type LoginServerLoginStart struct {
	Name string
}

type LoginServerEncryptionResponse struct {
	SharedSecretLen t.VarInt
	SharedSecret    []byte
	VerifyTokenLen  t.VarInt
	VerifyToken     []byte
}
