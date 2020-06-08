package util

import (
	"bytes"
	"encoding/binary"
)

func Uint64ToByte(n uint64) []byte {

	//todo
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, n)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
