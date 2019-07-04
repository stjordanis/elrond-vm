// File generated by the K Framework Go backend. Timestamp: 2019-07-04 01:26:11.488

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) freshFunction (s m.Sort, config m.KReference, counter int) (m.KReference, error) {
	switch s {
		case m.SortID:
			return i.evalFreshID(i.Model.FromInt(counter), config, -1)
		case m.SortInt:
			return i.evalFreshInt(i.Model.FromInt(counter), config, -1)
		default:
			panic("Cannot find fresh function for sort " + s.Name())
	}
}

