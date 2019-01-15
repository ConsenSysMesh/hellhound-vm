package hh

import "math/big"

//go:generate mockgen -destination=../mocks/mock_stack.go -package=mocks github.com/ConsenSys/hellhound-vm/hh Stack

type Stack interface {
	Data() []*big.Int
	Push(*big.Int)
	PushN(...*big.Int)
	Pop() (*big.Int)
	Peek() *big.Int
	Get(int) *big.Int
	Swap(int)
	Len() int
}