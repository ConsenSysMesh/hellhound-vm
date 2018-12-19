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
	configurers = append(configurers, Dispatcher())
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
	fmt.Println("****** stack ********")
	for i := vm.Stack().Len() - 1; i >= 0; i-- {
		if vm.Stack().Data()[i] != nil {
			fmt.Println(hex.EncodeToString(vm.Stack().Data()[i].Bytes()))
		}
	}
	fmt.Println("**************")
	fmt.Println("****** keystore ********")
	for slot, key := range vm.keystore.Keys() {
		if key != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, key.String())
		}
	}
	fmt.Println("**************")
	fmt.Println("****** register set ********")
	for slot, entry := range vm.registerSet.Values() {
		if entry != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, hex.EncodeToString(entry))
		}
	}
	fmt.Println("**************")
	fmt.Println("--------------------------------------------")

}
