package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
	"math/big"
)

func Stack() []hh.Operation {
	return []hh.Operation{
		hh.NewInstruction(hh.POP, pop),
		hh.NewInstruction(hh.PUSH, push),
	}
}

func pop(vm hh.VM, _ *hh.Contract) error{
	vm.Stack().Pop()
	return nil
}

func push(vm hh.VM, contract *hh.Contract) error{
	size := int(contract.Code[contract.GetAndMovePCForward()])
	value := big.NewInt(0).SetBytes(contract.Code[contract.PC() : contract.MovePCForwardN(size)])
	vm.Stack().Push(value)
	return nil
}