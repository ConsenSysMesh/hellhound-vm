package hh

import "math/big"

type Stack interface {
	Data() []*big.Int
	Push(*big.Int)
	PushN(...*big.Int)
	Pop() (*big.Int)
	Peek() *big.Int
	Get(n int) *big.Int
	Len() int
}