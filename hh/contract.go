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

func (contract Ki) PC() int {
	return contract.pc
}

func (contract *Ki) GetAndMovePCForward() int {
	hp := contract.pc
	contract.pc++
	return hp
}

func (contract *Ki) GetAndMovePCForwardN(n int) int {
	hp := contract.pc
	contract.pc += n
	return hp
}

func (contract *Ki) MovePCForward() int {
	contract.pc++
	return contract.pc
}

func (contract *Ki) MovePCForwardN(n int) int {
	contract.pc += n
	return contract.pc
}
