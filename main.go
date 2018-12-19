package main

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/core"
	"github.com/ConsenSys/hellhound-vm/hh"
	"log"
)

func main() {
	fmt.Println("starting Hellhound virtual machine")

	hhvm, err := core.NewVM(
		core.StackCfg(
			hh.StackSize(32),
		),
		core.KeystoreCfg(
			hh.KeystoreSlotNumber(8),
		),
		core.RegisterSetCfg(
			hh.RegisterSetSlotNumber(8),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HHVM version : ", hhvm.Version())

	code := []byte{
		byte(hh.LOADKEY), 0x05, 0x01, 0x00, 0x00, 0x04, 0x01, 0x02, 0x03, 0x04,
		byte(hh.LOADREG), 0x03, 0x00, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		byte(hh.PUSH), 0x03, 0xAB, 0XCD, 0XEF,
		byte(hh.PUSH), 0x02, 0x0A, 0X0B,
		byte(hh.PUSH), 0x04, 0x09, 0x08, 0x07, 0x06,
		byte(hh.PUSH), 0x01, 0x02,
		byte(hh.PUSH), 0x01, 0x01,
		byte(hh.ADD),
		byte(hh.PUSH), 0x01, 0x03,
		byte(hh.MUL),
		byte(hh.PUSH), 0x01, 0x01,
		byte(hh.SUB),
		byte(hh.PUSH), 0x01, 0x02,
		byte(hh.DIV),
	}

	hhContract := hh.NewContract(code)

	err = hhvm.Run(hhContract)
	hhvm.Dump()
	if err != nil {
		log.Fatal(err)
	}


}
