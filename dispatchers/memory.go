package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Memory() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.POPTOREG, popToReg),
	}
}

func popToReg(vm hh.Tanden, contract *hh.Ki) error{
	slot := int(contract.Kokyu[contract.GetAndMovePCForward()])
	value := vm.Stack().Pop()
	return vm.RegisterSet().Store(slot, value.Bytes())
}
