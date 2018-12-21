package core

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

var (
	defaultKeystoreOptions = hh.KeystoreOptions{
		SlotNumber: 1,
	}
)

type keystore struct {
	keys []*hh.Key
}

func KeystoreCfg(setters ...hh.KeystoreOption) TandenConfigurer {
	return func(_vm *tanden) {
		_vm.keystore = NewKeystore(setters...)
	}
}

func NewKeystore(setters ...hh.KeystoreOption) hh.Keystore {
	opts := &defaultKeystoreOptions
	for _, setter := range setters {
		setter(opts)
	}
	return KeystoreFromOptions(*opts)
}

func KeystoreFromOptions(opts hh.KeystoreOptions) hh.Keystore {
	keys := make([]*hh.Key, opts.SlotNumber)
	return &keystore{
		keys: keys,
	}
}

func (ks keystore) Store(slot int, key *hh.Key) error {
	if !ks.validSlot(slot) {
		return fmt.Errorf("invalid slot %d", slot)
	}
	if !key.IsValid() {
		return fmt.Errorf("invalid key")
	}
	ks.keys[slot] = key
	return nil
}

func (ks keystore) Get(slot int) (*hh.Key, error) {
	if !ks.validSlot(slot) {
		return nil, fmt.Errorf("invalid slot %d", slot)
	}
	key := ks.keys[slot]
	if key == nil {
		return nil, fmt.Errorf("empty slot")
	}
	return key, nil
}

func (ks keystore) Keys() []*hh.Key {
	return ks.keys
}

func (ks keystore) validSlot(slot int) bool {
	return slot < len(ks.keys)
}

