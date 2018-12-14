package core

import "github.com/ConsenSys/hellhound-vm/hh"

type dispatcher struct {
	instructions map[hh.Opcode]hh.Instruction
}

func DispatcherCfg(subDispatchers ...[]hh.OpCodeRoute) VMConfigurer{
	return func(_vm *vm) {
		_vm.dispatcher = NewDispatcher(subDispatchers...)
	}
}

func NewDispatcher(subDispatchers ...[]hh.OpCodeRoute) hh.Dispatcher {
	instructions := make(map[hh.Opcode]hh.Instruction)
	for _, routes := range subDispatchers {
		for _, route := range routes {
			instructions[route.Opcode] = route.Instruction
		}
	}
	return &dispatcher{
		instructions: instructions,
	}
}

func (dispatcher dispatcher) Dispatch(opcode hh.Opcode) (hh.Instruction, error) {
	instruction, exists := dispatcher.instructions[opcode]
	if !exists {
		return nil, InstructionNotFound{Opcode: opcode}
	}
	return instruction, nil
}
