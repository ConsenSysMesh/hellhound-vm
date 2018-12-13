package hh

type VM interface {
	Version() string
	Run([]byte) error
}



type Keystore interface {
	Store(slot int, key *Key) error
	Get(slot int) (*Key, error)
	Keys() []*Key
}

type RegisterSet interface {
	Store(slot int, entry []byte) error
	Get(slot int) ([]byte, error)
	Values() [][]byte
}

type Dispatcher interface {
	Dispatch(Opcode) (Instruction, error)
}

type Instruction func(*VM) error

type OpCodeRoute struct {
	Opcode Opcode
	Instruction    Instruction
}

func NewInstruction(opcode Opcode, instruction Instruction) OpCodeRoute {
	return OpCodeRoute{
		Opcode: opcode,
		Instruction:    instruction,
	}
}

type Key struct {
	Type  byte
	Usage byte
	Len   int
	Value []byte
}
