package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func RegisterSet() []hh.OpCodeRoute {
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADREG, LoadReg),
	}
}

func LoadReg(vm hh.VM) error {
	slot, value, err := LoadRegOperands(vm)
	if err != nil {
		return err
	}
	vm.RegisterSet().Store(int(slot), value)
	return nil
}

func LoadRegOperands(vm hh.VM) (slot byte, value []byte, err error) {
	slot = vm.Heap()[vm.GetAndMoveHPForward()]
	size := hh.GetLenInt(vm.Heap()[vm.HP():vm.MoveHPForwardN(2)])
	value = vm.Heap()[vm.HP():vm.MoveHPForwardN(size)]
	return
}
