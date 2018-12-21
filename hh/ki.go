package hh

type Ki struct {
	Code []byte
	// program counter
	pc int
}

func NewKi(code []byte) *Ki {
	return &Ki{
		Code: code,
		pc:   0,
	}
}

func (ki Ki) PC() int {
	return ki.pc
}

func (ki Ki) Burnable() bool{
	return !ki.Drained() && OpCode(ki.Code[ki.PC()]) != STOP
}

func (ki Ki) Drained() bool{
	return ki.PC() >= len(ki.Code)
}

func (ki *Ki) GetAndMovePCForward() int {
	hp := ki.pc
	ki.pc++
	return hp
}

func (ki *Ki) GetAndMovePCForwardN(n int) int {
	hp := ki.pc
	ki.pc += n
	return hp
}

func (ki *Ki) MovePCForward() int {
	ki.pc++
	return ki.pc
}

func (ki *Ki) MovePCForwardN(n int) int {
	ki.pc += n
	return ki.pc
}
