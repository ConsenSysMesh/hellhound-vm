package core

import (
	"errors"
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

var (
	defaultRegisterSetOptions = hh.RegisterSetOptions{
		SlotNumber: 1,
	}
)

type registerSet struct {
	values [][]byte
}

func RegisterSetCfg(setters ...hh.RegisterSetOption) VMConfigurer {
	return func(_vm *vm) {
		_vm.registerSet = NewRegisterSet(setters...)
	}
}

func NewRegisterSet(setters ...hh.RegisterSetOption) hh.RegisterSet {
	opts := &defaultRegisterSetOptions
	for _, setter := range setters {
		setter(opts)
	}
	return RegisterSetFromOptions(*opts)
}

func RegisterSetFromOptions(opts hh.RegisterSetOptions) hh.RegisterSet {
	return &registerSet{
		values: make([][]byte, opts.SlotNumber),
	}
}

func (rs registerSet) Store(slot int, entry []byte) error {
	if !rs.validSlot(slot) {
		return fmt.Errorf("invalid register set slot : %d", slot)
	}
	if entry == nil {
		return errors.New("cannot store empty registry entry")
	}
	rs.values[slot] = entry
	return nil
}

func (rs registerSet) Get(slot int) ([]byte, error) {
	if !rs.validSlot(slot) {
		return nil, fmt.Errorf("invalid register set slot : %d", slot)
	}
	return rs.values[slot], nil
}

func (rs registerSet) Values() [][]byte {
	return rs.values
}

func (rs registerSet) validSlot(slot int) bool {
	return slot > 0 && slot < len(rs.values)
}
