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
	stack       hh.Stack
	// stack pointer
	sp int
}

type VMConfigurer func(*vm)

func NewVM(configurers ...VMConfigurer) (hh.VM, error) {
	vm := &vm{}
	for _, configurer := range configurers {
		configurer(vm)
	}
	return vm, nil
}

func (vm *vm) Run(contract *hh.Contract) error {
	// reset VM values
	vm.reset()
	for contract.PC() < len(contract.Code) {
		// dispatch current opcode and get next instruction to execute
		instruction, err := vm.dispatcher.Dispatch(hh.OpCode(contract.Code[contract.GetAndMovePCForward()]))
		if err != nil {
			return err
		}
		// execute instruction
		err = instruction(vm, contract)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *vm) reset() {
	vm.sp = 0
}

func (vm vm) Stack() hh.Stack {
	return vm.stack
}

func (vm vm) SP() int {
	return vm.sp
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
	fmt.Println("\t STACK")
	for i, element := range vm.stack.Data() {
		if element != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", i, hex.EncodeToString(element.Bytes()))
		}
	}
	fmt.Println("\t KEYSTORE")
	fmt.Printf("\t\t slots : %d\n", len(vm.keystore.Keys()))
	for slot, key := range vm.keystore.Keys() {
		if key != nil {
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
