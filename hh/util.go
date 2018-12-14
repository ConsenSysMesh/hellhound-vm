package hh

import (
	"encoding/binary"
	"math/big"
)

func GetLenBytes(l int) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(l))
	return buf
}

func GetLenInt(b []byte) int {
	return int(b[1]) | int(b[0])<<8
}

func IntAsBytes(val int) []byte {
	p := new(big.Int)
	p.SetInt64(int64(val))
	return p.Bytes()
}
