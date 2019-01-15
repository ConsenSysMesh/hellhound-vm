package core

import "github.com/ConsenSys/hellhound-vm/hh"

type ki struct {
	kokyu hh.Kokyu
	// program counter
	pc int
}

func NewKi(kokyu hh.Kokyu) hh.Ki {
	return &ki{
		kokyu: kokyu,
		pc:    0,
	}
}


func (ki ki) Kokyu() []byte{
	return ki.kokyu
}

func (ki ki) PC() int {
	return ki.pc
}

func (ki ki) Burnable() bool{
	return !ki.Drained() && hh.OpCode(ki.Kokyu()[ki.PC()]) != hh.STOP
}

func (ki ki) Drained() bool{
	return ki.PC() > len(ki.Kokyu())
}

func (ki *ki) GetAndMovePCForward() int {
	hp := ki.pc
	ki.pc++
	return hp
}

func (ki *ki) GetAndMovePCForwardN(n int) int {
	hp := ki.pc
	ki.pc += n
	return hp
}

func (ki *ki) MovePCForward() int {
	ki.pc++
	return ki.pc
}

func (ki *ki) MovePCForwardN(n int) int {
	ki.pc += n
	return ki.pc
}

func (ki ki) ProvideHankoInput() []byte {
	return ki.Kokyu()
}

