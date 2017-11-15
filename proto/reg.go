package proto

var HandshakeServerIdToPacket map[int32]interface{}
var HandshakeServerPacketToId map[string]int32

var StatusServerIdToPacket map[int32]interface{}
var StatusServerPacketToId map[string]int32

var LoginServerIdToPacket map[int32]interface{}
var LoginServerPacketToId map[string]int32

var PlayServerIdToPacket map[int32]interface{}
var PlayServerPacketToId map[string]int32

func init() {
	HandshakeServerIdToPacket = make(map[int32]interface{})
	HandshakeServerIdToPacket[0x0] = HandshakeServerHandshake{}
	HandshakeServerIdToPacket[0xFE] = HandshakeServerHandshakeLegacy{}

	HandshakeServerPacketToId = make(map[string]int32)
	HandshakeServerPacketToId["proto.HandshakeServerHandshake"] = 0x00
	HandshakeServerPacketToId["proto.HandshakeServerHandshakeLegacy"] = 0xFE
}
