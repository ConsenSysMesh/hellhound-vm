package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func RegisterSet() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.LOADREG, loadReg),
	}
}

func loadReg(vm hh.Tanden, contract *hh.Ki) error {
	slot, value, err := loadRegOperands(contract)
	if err != nil {
		return err
	}
	vm.RegisterSet().Store(int(slot), value)
	return nil
}

func loadRegOperands(contract *hh.Ki) (slot byte, value []byte, err error) {
	slot = contract.Kokyu[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Kokyu[contract.PC():contract.MovePCForwardN(2)])
	value = contract.Kokyu[contract.PC():contract.MovePCForwardN(size)]
	return
}
