package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
	"math/big"
)

func Stack() []hh.Operation {
	return []hh.Operation{
		hh.NewInstruction(hh.POPN, pop),
		hh.NewInstruction(hh.PUSHN, push),
		hh.NewInstruction(hh.SWAPN, swap),
	}
}

func pop(vm hh.VM, contract *hh.Contract) error{
	size := int(contract.Code[contract.GetAndMovePCForward()])
	for i := 0; i < size; i++{
		vm.Stack().Pop()
	}
	return nil
}

func push(vm hh.VM, contract *hh.Contract) error{
	size := int(contract.Code[contract.GetAndMovePCForward()])
	value := big.NewInt(0).SetBytes(contract.Code[contract.PC() : contract.MovePCForwardN(size)])
	vm.Stack().Push(value)
	return nil
}

func swap(vm hh.VM, contract *hh.Contract) error{
	size := int(contract.Code[contract.GetAndMovePCForward()])
	vm.Stack().Swap(size)
	return nil
}