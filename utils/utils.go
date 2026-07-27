package utils

import "encoding/binary"

func IntToHex(num int64) []byte {
	return binary.BigEndian.AppendUint64(nil, uint64(num))
}
