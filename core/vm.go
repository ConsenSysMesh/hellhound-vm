package core

import (
	"encoding/hex"
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

const (
	version = "0.0.1"
)

type vm struct {
	keystore    hh.Keystore
	registerSet hh.RegisterSet
	dispatcher  hh.Dispatcher
	stack       []byte
	// stack pointer
	sp   int
	heap []byte
	// heap pointer
	hp int
}
type VMConfigurer func(*vm)

func NewVM(configurers ...VMConfigurer) (hh.VM, error) {
	vm := &vm{}
	for _, configurer := range configurers {
		configurer(vm)
	}
	return vm, nil
}

func (vm *vm) Run(code []byte) error {
	// reset VM values and copy code to the heap
	vm.reset(code)
	for vm.HP() < len(vm.Heap()) {
		// dispatch current opcode and get next instruction to execute
		instruction, err := vm.dispatcher.Dispatch(hh.Opcode(vm.Heap()[vm.GetAndMoveHPForward()]))
		if err != nil {
			return err
		}
		// execute instruction
		err = instruction(vm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *vm) reset(code []byte) {
	vm.heap = code
	vm.hp = 0
	vm.sp = 0
}

func (vm vm) Heap() []byte {
	return vm.heap
}

func (vm vm) Stack() []byte {
	return vm.stack
}

func (vm vm) HP() int {
	return vm.hp
}

func (vm vm) SP() int {
	return vm.sp
}

func (vm *vm) GetAndMoveHPForward() int {
	hp := vm.hp
	vm.hp++
	return hp
}

func (vm *vm) GetAndMoveHPForwardN(n int) int {
	hp := vm.hp
	vm.hp += n
	return hp
}

func (vm *vm) MoveHPForward() int {
	vm.hp++
	return vm.hp
}

func (vm *vm) MoveHPForwardN(n int) int {
	vm.hp += n
	return vm.hp
}

func (vm vm) Version() string {
	return version
}

func (vm vm) Keystore() hh.Keystore {
	return vm.keystore
}

func (vm vm) RegisterSet() hh.RegisterSet {
	return vm.registerSet
}

func (vm vm) Dump() {
	fmt.Println("--------------------------------------------")
	fmt.Println("VM DUMP")
	fmt.Println("\t KEYSTORE")
	fmt.Printf("\t\t slots : %d\n", len(vm.keystore.Keys()))
	for slot, key := range vm.keystore.Keys() {
		if key != nil{
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, key.String())
		}
	}
	fmt.Println("\t REGISTER SET")
	fmt.Printf("\t\t slots : %d\n", len(vm.registerSet.Values()))
	for slot, entry := range vm.registerSet.Values() {
		if entry != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, hex.EncodeToString(entry))
		}
	}
	fmt.Println("--------------------------------------------")

}
