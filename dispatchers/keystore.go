package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Keystore() []hh.OpCodeRoute{
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADKEY, LoadKey),
	}
}

func LoadKey(_ *hh.VM) error{

	return nil
}