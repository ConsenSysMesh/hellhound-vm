package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
)

func Keystore() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.LOADKEY, loadKey),
	}
}

func loadKey(vm hh.VM, contract *hh.Ki) error{
	slot, keyType, usage, keyValue, err := loadKeyOperands(contract)
	if err != nil{
		return err
	}
	key := hh.NewKey(keyType, usage, keyValue)
	return vm.Keystore().Store(int(slot), key)
}

func loadKeyOperands(contract *hh.Ki) (slot, keyType, usage byte, keyValue []byte, err error){
	slot = contract.Code[contract.GetAndMovePCForward()]
	keyType = contract.Code[contract.GetAndMovePCForward()]
	usage = contract.Code[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Code[contract.PC() : contract.MovePCForwardN(2)])
	keyValue = contract.Code[contract.PC() : contract.MovePCForwardN(size)]
	return
}