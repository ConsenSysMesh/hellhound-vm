package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
	"math/big"
)

func Stack() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.POPN, pop),
		hh.NewKiWave(hh.PUSHN, push),
		hh.NewKiWave(hh.SWAPN, swap),
	}
}

func pop(vm hh.Tanden, contract *hh.Ki) error{
	size := int(contract.Kokyu[contract.GetAndMovePCForward()])
	for i := 0; i < size; i++{
		vm.Stack().Pop()
	}
	return nil
}

func push(vm hh.Tanden, contract *hh.Ki) error{
	size := int(contract.Kokyu[contract.GetAndMovePCForward()])
	value := big.NewInt(0).SetBytes(contract.Kokyu[contract.PC() : contract.MovePCForwardN(size)])
	vm.Stack().Push(value)
	return nil
}

func swap(vm hh.Tanden, contract *hh.Ki) error{
	size := int(contract.Kokyu[contract.GetAndMovePCForward()])
	vm.Stack().Swap(size)
	return nil
}