package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func RegisterSet() []hh.Operation {
	return []hh.Operation{
		hh.NewInstruction(hh.LOADREG, loadReg),
	}
}

func loadReg(vm hh.VM, contract *hh.Contract) error {
	slot, value, err := loadRegOperands(contract)
	if err != nil {
		return err
	}
	vm.RegisterSet().Store(int(slot), value)
	return nil
}

func loadRegOperands(contract *hh.Contract) (slot byte, value []byte, err error) {
	slot = contract.Code[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Code[contract.PC():contract.MovePCForwardN(2)])
	value = contract.Code[contract.PC():contract.MovePCForwardN(size)]
	return
}
