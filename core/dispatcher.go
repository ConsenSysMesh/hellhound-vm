package core

import (
	"github.com/ConsenSys/hellhound-vm/dispatchers"
	"github.com/ConsenSys/hellhound-vm/hh"
)

type dispatcher struct {
	kiWaves map[hh.OpCode]hh.KiWave
}

func Dispatcher() TandenConfigurer {
	return DispatcherCfg(
		dispatchers.Arithmetic(),
		dispatchers.Stack(),
		dispatchers.Keystore(),
		dispatchers.RegisterSet(),
		dispatchers.Memory(),
		dispatchers.Paillier(),
	)
}

func DispatcherCfg(subDispatchers ...[]hh.Operation) TandenConfigurer {
	return func(_vm *tanden) {
		_vm.dispatcher = NewDispatcher(subDispatchers...)
	}
}

func NewDispatcher(subDispatchers ...[]hh.Operation) hh.Dispatcher {
	kiWaves := make(map[hh.OpCode]hh.KiWave)
	for _, operations := range subDispatchers {
		for _, route := range operations {
			kiWaves[route.OpCode] = route.KiWave
		}
	}
	return &dispatcher{
		kiWaves: kiWaves,
	}
}

func (dispatcher dispatcher) Dispatch(opcode hh.OpCode) (hh.KiWave, error) {
	instruction, exists := dispatcher.kiWaves[opcode]
	if !exists {
		return nil, KiWaveNotFound{OpCode: opcode}
	}
	return instruction, nil
}
