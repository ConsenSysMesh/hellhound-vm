package core

import "github.com/ConsenSys/hellhound-vm/hh"

var(
	defaultRegisterSetOptions = hh.RegisterSetOptions{
		SlotNumber: 1,
	}
)

type registerSet struct {

}

func RegisterSetCfg(setters ...hh.RegisterSetOption) VMConfigurer{
	return func(_vm *vm) {
		_vm.registerSet = NewRegisterSet(setters...)
	}
}

func NewRegisterSet(setters ...hh.RegisterSetOption) hh.RegisterSet{
	opts := &defaultRegisterSetOptions
	for _,setter := range setters{
		setter(opts)
	}
	return RegisterSetFromOptions(*opts)
}

func RegisterSetFromOptions(opts hh.RegisterSetOptions) hh.RegisterSet{
	return &registerSet{

	}
}

func (rs registerSet) Store(slot int, entry []byte) error {
	panic("implement me")
}

func (rs registerSet) Get(slot int) ([]byte, error) {
	panic("implement me")
}

func (rs registerSet) Values() [][]byte {
	panic("implement me")
}
