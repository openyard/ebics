// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type UserSignatureDataSigBookType struct {
	OrderSignature     *UserSignatureDataSigBookTypeOrderSignature `xml:"OrderSignature"`
	OrderSignatureData *OrderSignatureDataType                     `xml:"OrderSignatureData"`

	Any []*w3c.Any
}

func NewUserSignatureDataSigBookType() *UserSignatureDataSigBookType {
	return new(UserSignatureDataSigBookType)
}

func (me *UserSignatureDataSigBookType) SetOrderSignature(value *UserSignatureDataSigBookTypeOrderSignature) {
	me.OrderSignature = value
}

func (me *UserSignatureDataSigBookType) AddOrderSignature() *UserSignatureDataSigBookTypeOrderSignature {
	me.OrderSignature = new(UserSignatureDataSigBookTypeOrderSignature)
	return me.OrderSignature
}

func (me *UserSignatureDataSigBookType) SetOrderSignatureData(value *OrderSignatureDataType) {
	me.OrderSignatureData = value
}

func (me *UserSignatureDataSigBookType) AddOrderSignatureData() *OrderSignatureDataType {
	me.OrderSignatureData = new(OrderSignatureDataType)
	return me.OrderSignatureData
}
