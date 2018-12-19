package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
)

func Keystore() []hh.OpCodeRoute{
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADKEY, LoadKey),
	}
}

func LoadKey(vm hh.VM, contract *hh.Contract) error{
	slot, keyType, usage, keyValue, err := LoadKeyOperands(contract)
	if err != nil{
		return err
	}
	key := hh.NewKey(keyType, usage, keyValue)
	return vm.Keystore().Store(int(slot), key)
}

func LoadKeyOperands(contract *hh.Contract) (slot, keyType, usage byte, keyValue []byte, err error){
	slot = contract.Code[contract.GetAndMovePCForward()]
	keyType = contract.Code[contract.GetAndMovePCForward()]
	usage = contract.Code[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Code[contract.PC() : contract.MovePCForwardN(2)])
	keyValue = contract.Code[contract.PC() : contract.MovePCForwardN(size)]
	return
}