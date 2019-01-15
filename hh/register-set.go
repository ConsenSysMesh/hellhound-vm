package hh

//go:generate mockgen -destination=../mocks/mock_register-set.go -package=mocks github.com/ConsenSys/hellhound-vm/hh RegisterSet

type RegisterSet interface {
	HankoInputProvider
	Store(slot int, entry []byte) error
	Get(slot int) ([]byte, error)
	Values() [][]byte
}

