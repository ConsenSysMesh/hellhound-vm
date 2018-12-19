package core

import (
	"github.com/ConsenSys/hellhound-vm/hh"
	"math/big"
)

var (
	defaultStackOptions = hh.StackOptions{
		Size: 64,
	}
)

type stack struct {
	data []*big.Int
}

func StackCfg(setters ...hh.StackOption) VMConfigurer {
	return func(_vm *vm) {
		_vm.stack = NewStack(setters...)
	}
}

func NewStack(setters ...hh.StackOption) hh.Stack {
	opts := &defaultStackOptions
	for _, setter := range setters {
		setter(opts)
	}
	return StackFromOptions(*opts)
}

func StackFromOptions(opts hh.StackOptions) hh.Stack {
	return &stack{
		data: make([]*big.Int, 0, opts.Size),
	}
}

func (stack stack) Data() []*big.Int {
	return stack.data
}
