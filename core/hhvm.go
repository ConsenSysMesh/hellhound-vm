package core

import "github.com/ConsenSys/hellhound-vm/hh"

const (
	version = "0.0.1"
)

type vm struct {
	keystore    hh.Keystore
	registerSet hh.RegisterSet
	dispatcher  hh.Dispatcher
}

type VMConfigurer func(*vm)

func NewVM(configurers ...VMConfigurer) (hh.VM, error) {
	hhvm := &vm{}
	for _, configurer := range configurers{
		configurer(hhvm)
	}
	return hhvm, nil
}

func (hhvm vm) Run([]byte) error {

	return nil
}

func (hhvm vm) Version() string {
	return version
}
