package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
)

func Keystore() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.LOADKEY, loadKey),
	}
}

func loadKey(vm hh.Tanden, contract hh.Ki) error{
	slot, keyType, usage, keyValue, err := loadKeyOperands(contract)
	if err != nil{
		return err
	}
	key := hh.NewKey(keyType, usage, keyValue)
	return vm.Keystore().Store(int(slot), key)
}

func loadKeyOperands(contract hh.Ki) (slot, keyType, usage byte, keyValue []byte, err error){
	slot = contract.Kokyu()[contract.GetAndMovePCForward()]
	keyType = contract.Kokyu()[contract.GetAndMovePCForward()]
	usage = contract.Kokyu()[contract.GetAndMovePCForward()]
	size := hh.GetLenInt(contract.Kokyu()[contract.PC() : contract.MovePCForwardN(2)])
	keyValue = contract.Kokyu()[contract.PC() : contract.MovePCForwardN(size)]
	return
}