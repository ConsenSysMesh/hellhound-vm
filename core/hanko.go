package core

import (
	"github.com/ConsenSys/hellhound-vm/hh"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

var(
	hankoSensei = sha3HankoSensei{}
)

type sha3HankoSensei struct {

}

func HankoSensei() hh.HankoSensei{
	return hankoSensei
}

func (sha3HankoSensei) Hanko(input hh.HankoInputProvider) []byte {
	return hanko(input.ProvideHankoInput())
}

func (sha3HankoSensei) HankoN(hankoInputProviders ...hh.HankoInputProvider) []byte {
	var buf []byte
	for _, hankoInputProvider := range hankoInputProviders{
		buf = append(buf, hankoInputProvider.ProvideHankoInput()...)
	}
	return hanko(buf)
}

func hanko(input []byte) []byte{
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write(input)
	buf = hash.Sum(buf)
	return buf
}