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
	Burn(*Ki) error
}

type KiSensei interface {
	Flows(*Ki) (KiWave, error)
}

type Keystore interface {
	HankoInputProvider
	Store(slot int, key *Key) error
	Get(slot int) (*Key, error)
	Keys() []*Key
}

type RegisterSet interface {
	HankoInputProvider
	Store(slot int, entry []byte) error
	Get(slot int) ([]byte, error)
	Values() [][]byte
}

type Dispatcher interface {
	Dispatch(OpCode) (KiWave, error)
}

type KiWave func(Tanden,*Ki) error

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

