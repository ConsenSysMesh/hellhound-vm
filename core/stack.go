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

func (stack stack) Data() []*big.Int {
	return stack.data
}

func (stack *stack) Push(d *big.Int) {
	stack.data = append(stack.data, d)
}

func (stack *stack) PushN(ds ...*big.Int) {
	stack.data = append(stack.data, ds...)
}

func (stack *stack) Pop() (ret *big.Int) {
	ret = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return
}

func (stack stack) Len() int {
	return len(stack.data)
}

func (stack stack) Peek() *big.Int {
	return stack.data[stack.Len()-1]
}

func (stack *stack) Get(n int) *big.Int {
	return stack.data[stack.Len()-n-1]
}

func (stack *stack) Swap(n int) {
	stack.data[stack.Len()-n], stack.data[stack.Len()-1] = stack.data[stack.Len()-1], stack.data[stack.Len()-n]
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
