// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
