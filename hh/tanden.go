package hh

import (
	"encoding/hex"
)

type Tanden interface {
	KiBurner
	KiSensei
	HankoInputProviders() []HankoInputProvider
	Keystore() Keystore
	RegisterSet() RegisterSet
	Stack() Stack
	Version() string
	Dump()
}

type KiBurner interface {
	Burn(Ki) error
}

type KiSensei interface {
	Flows(Ki) (KiWave, error)
}

type KiWave func(Tanden,Ki) error

type Operation struct {
	OpCode      OpCode
	KiWave KiWave
}

func NewKiWave(opcode OpCode, kiWave KiWave) Operation {
	return Operation{
		OpCode:      opcode,
		KiWave: kiWave,
	}
}

type Key struct {
	Type  byte
	Usage byte
	Len   int
	Value []byte
}

func NewKey(keyType, keyUsage byte, value []byte) *Key {
	return &Key{
		Type:  keyType,
		Usage: keyUsage,
		Len:   len(value),
		Value: value,
	}
}

func (k Key) String() string {
	return hex.EncodeToString(k.Value)
}

func (k Key) IsValid() bool {
	return k.Len == len(k.Value)
}

