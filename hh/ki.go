package hh

type Kokyu []byte

//go:generate mockgen -destination=../mocks/mock_ki.go -package=mocks github.com/ConsenSys/hellhound-vm/hh Ki

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
