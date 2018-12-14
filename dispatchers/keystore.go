package dispatchers

import (
	"github.com/ConsenSys/hellhound-vm/hh"
)

func Keystore() []hh.OpCodeRoute{
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.LOADKEY, LoadKey),
	}
}

func LoadKey(vm hh.VM) error{
	slot, keyType, usage, keyValue, err := LoadKeyOperands(vm)
	if err != nil{
		return err
	}
	key := hh.NewKey(keyType, usage, keyValue)
	return vm.Keystore().Store(int(slot), key)
}

func LoadKeyOperands(vm hh.VM) (slot, keyType, usage byte, keyValue []byte, err error){
	slot = vm.Heap()[vm.GetAndMoveHPForward()]
	keyType = vm.Heap()[vm.GetAndMoveHPForward()]
	usage = vm.Heap()[vm.GetAndMoveHPForward()]
	size := hh.GetLenInt(vm.Heap()[vm.HP() : vm.MoveHPForwardN(2)])
	keyValue = vm.Heap()[vm.HP() : vm.MoveHPForwardN(size)]
	return
}