package hh

//go:generate mockgen -destination=../mocks/mock_hanko_input_provider.go -package=mocks github.com/ConsenSys/hellhound-vm/hh HankoInputProvider
//go:generate mockgen -destination=../mocks/mock_hanko.go -package=mocks github.com/ConsenSys/hellhound-vm/hh HankoSensei


type HankoInputProvider interface {
	ProvideHankoInput() []byte
}

type HankoSensei interface {
	Hanko(HankoInputProvider) []byte
	HankoN(...HankoInputProvider) []byte
}
