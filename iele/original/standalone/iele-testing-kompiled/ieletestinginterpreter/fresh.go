// File generated by the K Framework Go backend. Timestamp: 2019-07-15 11:26:24.140

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
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

