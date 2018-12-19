package core

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

type InstructionNotFound struct {
	Opcode hh.OpCode
}

func (inf InstructionNotFound) Error() string {
	return fmt.Sprintf("instruction not found for opcode : %X", inf.Opcode)
}

