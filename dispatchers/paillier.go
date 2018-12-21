package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Paillier() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.PAILLIERADDCIPHERS, paillierAddCiphers),
		hh.NewKiWave(hh.PAILIERADDCONSTANT, paillierAddConstant),
		hh.NewKiWave(hh.PAILIERMULCONSTANT, paillierMulConstant),
	}
}

func paillierAddCiphers(_ hh.VM, _ *hh.Ki) error {

	return nil
}

func paillierAddConstant(_ hh.VM, _ *hh.Ki) error {

	return nil
}

func paillierMulConstant(_ hh.VM, _ *hh.Ki) error {

	return nil
}
