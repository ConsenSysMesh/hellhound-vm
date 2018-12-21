package core

import (
	"encoding/hex"
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

const (
	version = "0.0.1"
)

type tanden struct {
	keystore    hh.Keystore
	registerSet hh.RegisterSet
	dispatcher  hh.Dispatcher
	stack       hh.Stack
	// stack pointer
	sp int
}

type TandenConfigurer func(*tanden)

func NewTanden(configurers ...TandenConfigurer) (hh.Tanden, error) {
	tanden := &tanden{}
	configurers = append(configurers, Dispatcher())
	for _, configurer := range configurers {
		configurer(tanden)
	}
	return tanden, nil
}

func (tanden *tanden) Burn(ki *hh.Ki) error {
	// reset Tanden values
	tanden.reset()
	// check if Ki is burnable, meaning the current OpCode is not STOP
	for ki.Burnable() {
		// dispatch current opcode and get next instruction to execute
		// the Ki flows
		wave, err := tanden.Flows(ki)
		if err != nil {
			return err
		}
		// execute instruction
		err = wave(tanden, ki)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tanden *tanden) Flows(ki *hh.Ki) (hh.KiWave, error) {
	return tanden.dispatcher.Dispatch(hh.OpCode(ki.Kokyu[ki.GetAndMovePCForward()]))
}

func (tanden *tanden) reset() {
	tanden.sp = 0
}

func (tanden tanden) Stack() hh.Stack {
	return tanden.stack
}

func (tanden tanden) SP() int {
	return tanden.sp
}

func (tanden tanden) Version() string {
	return version
}

func (tanden tanden) Keystore() hh.Keystore {
	return tanden.keystore
}

func (tanden tanden) RegisterSet() hh.RegisterSet {
	return tanden.registerSet
}

func (tanden tanden) HankoInputProviders() []hh.HankoInputProvider {
	return []hh.HankoInputProvider{
		tanden.keystore,
		tanden.registerSet,
	}
}

func (tanden tanden) Dump() {
	fmt.Println("--------------------------------------------")
	fmt.Println("****** stack ********")
	for i := tanden.Stack().Len() - 1; i >= 0; i-- {
		if tanden.Stack().Data()[i] != nil {
			fmt.Println(hex.EncodeToString(tanden.Stack().Data()[i].Bytes()))
		}
	}
	fmt.Println("**************")
	fmt.Println("****** keystore ********")
	for slot, key := range tanden.keystore.Keys() {
		if key != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, key.String())
		}
	}
	fmt.Println("**************")
	fmt.Println("****** register set ********")
	for slot, entry := range tanden.registerSet.Values() {
		if entry != nil {
			fmt.Printf("\t\t [ %3d ] = %s\n", slot, hex.EncodeToString(entry))
		}
	}
	fmt.Println("**************")
	fmt.Println("--------------------------------------------")

}
