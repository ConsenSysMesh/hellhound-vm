package hh

type OpCode byte

// 0x0 range - arithmetic ops.
const(
	STOP OpCode = iota
	ADD
	SUB
	MUL
	DIV
)

// 0x10 range - comparison ops
const(

)

// 0x20 range - storage and execution
const(
	POPN OpCode = iota + 0x20
	PUSHN
	SWAPN
)

// 0x30 range - generic vm components ops.
const(
	LOADKEY OpCode = iota + 0x30
	LOADREG
)

// 0x40 range - Paillier ops.
const(
	PAILLIERADDCIPHERS OpCode = iota + 0x40
	PAILIERADDCONSTANT
	PAILIERMULCONSTANT
)
