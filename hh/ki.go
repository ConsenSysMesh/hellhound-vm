package hh

type Kokyu []byte

type Ki struct {
	Kokyu Kokyu
	// program counter
	pc int
}

func NewKi(kokyu Kokyu) *Ki {
	return &Ki{
		Kokyu: kokyu,
		pc:    0,
	}
}

func (ki Ki) PC() int {
	return ki.pc
}

func (ki Ki) Burnable() bool{
	return !ki.Drained() && OpCode(ki.Kokyu[ki.PC()]) != STOP
}

func (ki Ki) Drained() bool{
	return ki.PC() >= len(ki.Kokyu)
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
