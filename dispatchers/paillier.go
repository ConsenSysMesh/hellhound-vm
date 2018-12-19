package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Paillier() []hh.Operation {
	return []hh.Operation{
		hh.NewInstruction(hh.PAILLIERADDCIPHERS, paillierAddCiphers),
		hh.NewInstruction(hh.PAILIERADDCONSTANT, paillierAddConstant),
		hh.NewInstruction(hh.PAILIERMULCONSTANT, paillierMulConstant),
	}
}

func paillierAddCiphers(_ hh.VM, _ *hh.Contract) error {

	return nil
}

func paillierAddConstant(_ hh.VM, _ *hh.Contract) error {

	return nil
}

func paillierMulConstant(_ hh.VM, _ *hh.Contract) error {

	return nil
}
