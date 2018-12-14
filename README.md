# Hellhound Virtual Machine

Hellhound Virtual Machine (HHVM) is a byte code virtual machine.

## Structure

The HHVM is compose of the following elements :

- Register Set
- Keystore
- Dispatcher

## Initialize VM

```go
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
```