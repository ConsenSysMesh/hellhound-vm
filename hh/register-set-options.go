package hh

type RegisterSetOptions struct {
	SlotNumber int
}

type RegisterSetOption func(*RegisterSetOptions)

func RegisterSetSlotNumber(slotNumber int) RegisterSetOption {
	return func(options *RegisterSetOptions) {
		options.SlotNumber = slotNumber
	}
}
