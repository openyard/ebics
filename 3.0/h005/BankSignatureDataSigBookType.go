// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type BankSignatureDataSigBookType struct {
	OrderSignature *BankSignatureDataSigBookTypeOrderSignature `xml:"OrderSignature,omitempty"`

	Any []*w3c.Any
}

func NewBankSignatureDataSigBookType() *BankSignatureDataSigBookType {
	return new(BankSignatureDataSigBookType)
}

func (me *BankSignatureDataSigBookType) SetOrderSignature(value *BankSignatureDataSigBookTypeOrderSignature) {
	me.OrderSignature = value
}

func (me *BankSignatureDataSigBookType) AddOrderSignature() *BankSignatureDataSigBookTypeOrderSignature {
	me.OrderSignature = new(BankSignatureDataSigBookTypeOrderSignature)
	return me.OrderSignature
}
