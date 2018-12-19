package main

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/core"
	"github.com/ConsenSys/hellhound-vm/dispatchers"
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
		core.DispatcherCfg(
			dispatchers.Keystore(),
			dispatchers.RegisterSet(),
			dispatchers.Paillier(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HHVM version : ", hhvm.Version())

	code := []byte{
		byte(hh.LOADKEY), 0x05, 0x01, 0x00, 0x00, 0x04, 0x01, 0x02, 0x03, 0x04,
		byte(hh.LOADREG), 0x03, 0x00, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
	}

	hhContract := hh.NewContract(code)

	err = hhvm.Run(hhContract)

	if err != nil {
		log.Fatal(err)
	}

	hhvm.Dump()
}
