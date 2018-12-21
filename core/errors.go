package core

import (
	"fmt"
	"github.com/ConsenSys/hellhound-vm/hh"
)

type KiWaveNotFound struct {
	OpCode hh.OpCode
}

func (inf KiWaveNotFound) Error() string {
	return fmt.Sprintf("instruction not found for opcode : %X", inf.OpCode)
}

