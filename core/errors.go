package core

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

type InstructionNotFound struct {
	Opcode hh.Opcode
}

func (inf InstructionNotFound) Error() string {
	return fmt.Sprintf("instruction not found for opcode : %X", inf.Opcode)
}

