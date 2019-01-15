package hh

//go:generate mockgen -destination=../mocks/mock_keystore.go -package=mocks github.com/ConsenSys/hellhound-vm/hh Keystore

type Keystore interface {
	HankoInputProvider
	Store(slot int, key *Key) error
	Get(slot int) (*Key, error)
	Keys() []*Key
}
