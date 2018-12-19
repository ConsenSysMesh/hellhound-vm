package hh

type OpCode byte

// 0x0 range - arithmetic ops.
const(
	STOP OpCode = iota
)

// 0x10 range - comparison ops
const(

)

// 0x20 range - generic vm components ops.
const(
	LOADKEY OpCode = iota + 0x20
	LOADREG
)

// 0x30 range - Paillier ops.
const(
	PAILLIERADDCIPHERS OpCode = iota + 0x30
	PAILIERADDCONSTANT
	PAILIERMULCONSTANT
)
