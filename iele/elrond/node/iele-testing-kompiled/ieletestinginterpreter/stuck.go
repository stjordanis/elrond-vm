// File generated by the K Framework Go backend. Timestamp: 2019-07-30 16:33:19.058

package ieletestinginterpreter 

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) makeStuck(c m.KReference, config m.KReference) (m.KReference, error) {
	matched := false
	// rule #-1
	// source: ? @?
	// {| rule `<generatedTop>`(_0,`<s>`(``.K=>#STUCK(.KList)``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token("true","Bool") ensures #token("true","Bool") [] |}
	if !matched {
		// LHS
		if c&kapplyMatchMask == kapplyMatchLblXltgeneratedTopXgt9 { // `<generatedTop>`(_0,`<s>`(DotVar1),_1,_2,_3,_4,_5,_6,_7)
			kapp0 := c
			varXu0 := i.Model.KApplyArg(kapp0, 0) // lhs KVariable _0
			if i.Model.KApplyArg(kapp0, 1)&kapplyMatchMask == kapplyMatchLblXltsXgt1 { // `<s>`(DotVar1)
				kapp1 := i.Model.KApplyArg(kapp0, 1)
				// KSequence, size 1:DotVar1
				varDotVar1 := i.Model.KApplyArg(kapp1, 0) // lhs KVariable DotVar1
				varXu1 := i.Model.KApplyArg(kapp0, 2) // lhs KVariable _1
				varXu2 := i.Model.KApplyArg(kapp0, 3) // lhs KVariable _2
				varXu3 := i.Model.KApplyArg(kapp0, 4) // lhs KVariable _3
				varXu4 := i.Model.KApplyArg(kapp0, 5) // lhs KVariable _4
				varXu5 := i.Model.KApplyArg(kapp0, 6) // lhs KVariable _5
				varXu6 := i.Model.KApplyArg(kapp0, 7) // lhs KVariable _6
				varXu7 := i.Model.KApplyArg(kapp0, 8) // lhs KVariable _7
				// RHS
				i.traceRuleApply("STEP", -1, "{| rule `<generatedTop>`(_0,`<s>`(``.K=>#STUCK(.KList)``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token(\"true\",\"Bool\") ensures #token(\"true\",\"Bool\") [] |}")
				return i.Model.NewKApply(m.LblXltgeneratedTopXgt, // as-is <generatedTop>
					varXu0,
					i.Model.NewKApply(m.LblXltsXgt, // as-is <s>
						i.Model.AssembleKSequence(
							i.Model.NewKApply(m.LblXhashSTUCK, // as-is #STUCK
							),
							varDotVar1,
						),
					),
					varXu1,
					varXu2,
					varXu3,
					varXu4,
					varXu5,
					varXu6,
					varXu7,
				), nil
			}
		}
	}

	return c, nil
}

func (i *Interpreter) makeUnstuck(c m.KReference, config m.KReference) (m.KReference, error) {
	matched := false
	// rule #-1
	// source: ? @?
	// {| rule `<generatedTop>`(_0,`<s>`(``#STUCK(.KList)=>.K``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token("true","Bool") ensures #token("true","Bool") [] |}
	if !matched {
		// LHS
		if c&kapplyMatchMask == kapplyMatchLblXltgeneratedTopXgt9 { // `<generatedTop>`(_0,`<s>`(#STUCK(.KList)~>DotVar1),_1,_2,_3,_4,_5,_6,_7)
			kapp0 := c
			varXu0 := i.Model.KApplyArg(kapp0, 0) // lhs KVariable _0
			if i.Model.KApplyArg(kapp0, 1)&kapplyMatchMask == kapplyMatchLblXltsXgt1 { // `<s>`(#STUCK(.KList)~>DotVar1)
				kapp1 := i.Model.KApplyArg(kapp0, 1)
				if kseq2 := i.Model.KApplyArg(kapp1, 0); kseq2>>refTypeShift != refEmptyKseqTypeAsUint { // #STUCK(.KList)~>DotVar1
					_, kseq2Head0, kseq2Tail0 := i.Model.KSequenceSplitHeadTail(kseq2) // #STUCK(.KList) ~> ...
					if kseq2Head0&kapplyMatchMask == kapplyMatchLblXhashSTUCK0 { // #STUCK(.KList)
						varDotVar1 := kseq2Tail0 // lhs KVariable DotVar1
						varXu1 := i.Model.KApplyArg(kapp0, 2) // lhs KVariable _1
						varXu2 := i.Model.KApplyArg(kapp0, 3) // lhs KVariable _2
						varXu3 := i.Model.KApplyArg(kapp0, 4) // lhs KVariable _3
						varXu4 := i.Model.KApplyArg(kapp0, 5) // lhs KVariable _4
						varXu5 := i.Model.KApplyArg(kapp0, 6) // lhs KVariable _5
						varXu6 := i.Model.KApplyArg(kapp0, 7) // lhs KVariable _6
						varXu7 := i.Model.KApplyArg(kapp0, 8) // lhs KVariable _7
						// RHS
						i.traceRuleApply("STEP", -1, "{| rule `<generatedTop>`(_0,`<s>`(``#STUCK(.KList)=>.K``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token(\"true\",\"Bool\") ensures #token(\"true\",\"Bool\") [] |}")
						return i.Model.NewKApply(m.LblXltgeneratedTopXgt, // as-is <generatedTop>
							varXu0,
							i.Model.NewKApply(m.LblXltsXgt, // as-is <s>/* rhs KSequence size=1 */ 
								varDotVar1,
							),
							varXu1,
							varXu2,
							varXu3,
							varXu4,
							varXu5,
							varXu6,
							varXu7,
						), nil
					}
				}
			}
		}
	}

	return c, nil
}

