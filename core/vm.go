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

func (vm *vm) Burn(ki *hh.Ki) error {
	// reset VM values
	vm.reset()
	// check if Ki is burnable, meaning the current OpCode is not STOP
	for ki.Burnable(){
		// dispatch current opcode and get next instruction to execute
		wave, err := vm.Flows(ki)
		if err != nil {
			return err
		}
		// execute instruction
		err = wave(vm, ki)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *vm) Flows(ki *hh.Ki) (hh.KiWave, error){
	return vm.dispatcher.Dispatch(hh.OpCode(ki.Code[ki.GetAndMovePCForward()]))
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
