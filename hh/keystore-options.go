package hh

type KeystoreOptions struct{
	SlotNumber int
}

type KeystoreOption func(*KeystoreOptions)

func KeystoreSlotNumber(slotNumber int) KeystoreOption{
	return func(opts *KeystoreOptions) {
		opts.SlotNumber = slotNumber
	}
}