# Hellhound Virtual Machine

Hellhound Virtual Machine (HHVM) is a byte code virtual machine.

## Structure

The HHVM is compose of the following elements :

- Register Set
- Keystore
- Dispatcher

## Initialize VM

```go
	fmt.Println("creating Tanden virtual machine")
	tanden, err := core.NewTanden(
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

	fmt.Println("HHVM version : ", tanden.Version())

	kokyu := hh.Kokyu{
		byte(hh.LOADKEY), 0x05, 0x01, 0x00, 0x00, 0x04, 0x01, 0x02, 0x03, 0x04,
		byte(hh.LOADREG), 0x03, 0x00, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		byte(hh.PUSHN), 0x01, 0x02,
		byte(hh.PUSHN), 0x01, 0x01,
		byte(hh.ADD),
		byte(hh.PUSHN), 0x01, 0x03,
		byte(hh.MUL),
		byte(hh.PUSHN), 0x01, 0x01,
		byte(hh.SUB),
		byte(hh.PUSHN), 0x01, 0x02,
		byte(hh.SWAPN), 0x02,
		byte(hh.DIV),
		byte(hh.POPTOREG), 0x01,
		byte(hh.STOP),
	}

	ki := hh.NewKi(kokyu)

	err = tanden.Burn(ki)

	tanden.Dump()
```

