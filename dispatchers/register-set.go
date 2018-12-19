package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func RegisterSet() []hh.OpCodeRoute {
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADREG, LoadReg),
	}
}

func LoadReg(vm hh.VM, contract *hh.Contract) error {
	slot, value, err := LoadRegOperands(contract)
	if err != nil {
		return err
	}
	vm.RegisterSet().Store(int(slot), value)
	return nil
}

func LoadRegOperands(contract *hh.Contract) (slot byte, value []byte, err error) {
	slot = contract.Code[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Code[contract.PC():contract.MovePCForwardN(2)])
	value = contract.Code[contract.PC():contract.MovePCForwardN(size)]
	return
}
