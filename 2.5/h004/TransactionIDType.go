// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type TransactionIDType w3c.HexBinary

func NewTransactionIDType(value *w3c.HexBinary) *TransactionIDType {
	me := (*TransactionIDType)(value)
	return me
}
