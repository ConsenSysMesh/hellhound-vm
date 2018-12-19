package core

import (
	"github.com/ConsenSys/hellhound-vm/dispatchers"
	"github.com/ConsenSys/hellhound-vm/hh"
)

type dispatcher struct {
	instructions map[hh.OpCode]hh.Instruction
}

func Dispatcher() VMConfigurer{
	return DispatcherCfg(
		dispatchers.Arithmetic(),
		dispatchers.Stack(),
		dispatchers.Keystore(),
		dispatchers.RegisterSet(),
		dispatchers.Paillier(),
	)
}

func DispatcherCfg(subDispatchers ...[]hh.Operation) VMConfigurer{
	return func(_vm *vm) {
		_vm.dispatcher = NewDispatcher(subDispatchers...)
	}
}

func NewDispatcher(subDispatchers ...[]hh.Operation) hh.Dispatcher {
	instructions := make(map[hh.OpCode]hh.Instruction)
	for _, operations := range subDispatchers {
		for _, route := range operations {
			instructions[route.OpCode] = route.Instruction
		}
	}
	return &dispatcher{
		instructions: instructions,
	}
}

func (dispatcher dispatcher) Dispatch(opcode hh.OpCode) (hh.Instruction, error) {
	instruction, exists := dispatcher.instructions[opcode]
	if !exists {
		return nil, InstructionNotFound{Opcode: opcode}
	}
	return instruction, nil
}
