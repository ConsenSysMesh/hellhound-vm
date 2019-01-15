package hh

//go:generate mockgen -destination=../mocks/mock_dispatcher.go -package=mocks github.com/ConsenSys/hellhound-vm/hh Dispatcher

type Dispatcher interface {
	Dispatch(OpCode) (KiWave, error)
}
