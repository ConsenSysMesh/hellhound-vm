package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Paillier() []hh.OpCodeRoute {
	return []hh.OpCodeRoute{
		hh.NewInstruction(hh.PAILLIERADDCIPHERS, PaillierAddCiphers),
		hh.NewInstruction(hh.PAILIERADDCONSTANT, PaillierAddConstant),
		hh.NewInstruction(hh.PAILIERMULCONSTANT, PaillierMulConstant),
	}
}

func PaillierAddCiphers(_ hh.VM, _ *hh.Contract) error {

	return nil
}

func PaillierAddConstant(_ hh.VM, _ *hh.Contract) error {

	return nil
}

func PaillierMulConstant(_ hh.VM, _ *hh.Contract) error {

	return nil
}
