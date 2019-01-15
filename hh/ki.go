package hh

type Kokyu []byte

type Ki interface {
	Kokyu() []byte
	PC() int
	Burnable() bool
	Drained() bool
	GetAndMovePCForward() int
	GetAndMovePCForwardN(n int) int
	MovePCForward() int
	MovePCForwardN(n int) int
	HankoInputProvider
}
