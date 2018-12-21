package hh

import (
	"encoding/hex"
)

type VM interface {
	Keystore() Keystore
	RegisterSet() RegisterSet
	Stack() Stack
	Version() string
	Consume(*Ki) error
	Dump()
}

type Keystore interface {
	Store(slot int, key *Key) error
	Get(slot int) (*Key, error)
	Keys() []*Key
}

type RegisterSet interface {
	Store(slot int, entry []byte) error
	Get(slot int) ([]byte, error)
	Values() [][]byte
}

type Dispatcher interface {
	Dispatch(OpCode) (Instruction, error)
}

type Instruction func(VM,*Ki) error

type Operation struct {
	OpCode      OpCode
	Instruction Instruction
}

func NewInstruction(opcode OpCode, instruction Instruction) Operation {
	return Operation{
		OpCode:      opcode,
		Instruction: instruction,
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

