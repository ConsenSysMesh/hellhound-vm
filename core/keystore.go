package core

import "github.com/ConsenSys/hellhound-vm/hh"

var(
	defaultKeystoreOptions = hh.KeystoreOptions{
		SlotNumber: 1,
	}
)

type keystore struct {

}

func KeystoreCfg(setters ...hh.KeystoreOption) VMConfigurer{
	return func(_vm *vm) {
		_vm.keystore = NewKeystore(setters...)
	}
}

func NewKeystore(setters ...hh.KeystoreOption) hh.Keystore{
	opts := &defaultKeystoreOptions
	for _, setter := range setters{
		setter(opts)
	}
	return KeystoreFromOptions(*opts)
}

func KeystoreFromOptions(opts hh.KeystoreOptions) hh.Keystore{
	return &keystore{

	}
}

func (ks keystore) Store(slot int, key *hh.Key) error {
	panic("implement me")
}

func (ks keystore) Get(slot int) (*hh.Key, error) {
	panic("implement me")
}

func (ks keystore) Keys() []*hh.Key {
	panic("implement me")
}

