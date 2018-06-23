package password

import "encoding/binary"

func writeUint32BE(data []byte, v uint32) bool {
	binary.BigEndian.PutUint32(data, v)
	return true
}

func readUint32BE(data []byte, v uint32) uint32 {
	r := data[v:]
	return binary.BigEndian.Uint32(r)
}
