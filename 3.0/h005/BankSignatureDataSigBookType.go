// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
