package hh

type HankoInputProvider interface {
	ProvideHankoInput() []byte
}

type HankoSensei interface {
	Hanko(HankoInputProvider) []byte
	HankoN(...HankoInputProvider) []byte
}
