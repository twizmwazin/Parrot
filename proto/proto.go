package proto

type Proto int32

const (
	ProtoHandshake Proto = iota
	ProtoStatus
	ProtoLogin
	ProtoPlay
)
