package hh

type Opcode byte

const(
	STOP Opcode = 0x00

	// General purpose opcodes
	LOADKEY Opcode = 0x20
	LOADREG Opcode = 0x21

	// Paillier opcodes
	PAILLIERADDCIPHERS Opcode = 0x30
	PAILIERADDCONSTANT Opcode = 0x31
	PAILIERMULCONSTANT Opcode = 0x32
)
