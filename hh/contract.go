package hh

type Contract struct {
	Code []byte
	// program counter
	pc int
}

func NewContract(code []byte) *Contract {
	return &Contract{
		Code: code,
		pc:   0,
	}
}

func (contract Contract) PC() int {
	return contract.pc
}

func (contract *Contract) GetAndMovePCForward() int {
	hp := contract.pc
	contract.pc++
	return hp
}

func (contract *Contract) GetAndMovePCForwardN(n int) int {
	hp := contract.pc
	contract.pc += n
	return hp
}

func (contract *Contract) MovePCForward() int {
	contract.pc++
	return contract.pc
}

func (contract *Contract) MovePCForwardN(n int) int {
	contract.pc += n
	return contract.pc
}
