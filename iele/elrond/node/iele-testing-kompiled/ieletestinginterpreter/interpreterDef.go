// File generated by the K Framework Go backend. Timestamp: 2019-07-04 01:26:11.488

package ieletestinginterpreter

import (
	blockchain "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/blockchain"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/krypto"
)

// Interpreter is a container with a reference to model and basic options
type Interpreter struct {
	Model         *m.ModelState
	currentStep   int
	MaxSteps      int
	state         m.KReference
	traceHandlers []traceHandler
	Verbose       bool

	blockchainRef *blockchain.Blockchain
	kryptoRef *krypto.Krypto
}
// NewInterpreter creates a new interpreter instance
func NewInterpreter(blockchainRef *blockchain.Blockchain, kryptoRef *krypto.Krypto) *Interpreter {
	model := &m.ModelState{}
	model.Init()

	return &Interpreter {
		Model:         model,
		MaxSteps:      0,
		currentStep:   -1, // meaning that no processing started yet
		state:         m.NullReference,
		traceHandlers: nil,
		Verbose:       false,
		blockchainRef: blockchainRef,
		kryptoRef: kryptoRef,
	}
}
