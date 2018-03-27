// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type UserSignatureDataSigBookType struct {
	OrderSignatureData []*OrderSignatureDataType `xml:"OrderSignatureData"`

	Any []*w3c.Any
}

func NewUserSignatureDataSigBookType() *UserSignatureDataSigBookType {
	return new(UserSignatureDataSigBookType)
}

func (me *UserSignatureDataSigBookType) AddOrderSignatureData(value *OrderSignatureDataType) {
	me.OrderSignatureData = append(me.OrderSignatureData, value)
}
