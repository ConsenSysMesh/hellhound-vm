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
}
