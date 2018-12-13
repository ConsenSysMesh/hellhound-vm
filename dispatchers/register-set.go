package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func RegisterSet() []hh.OpCodeRoute{
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADREG, LoadReg),
	}
}

func LoadReg(_ *hh.VM) error{

	return nil
}
