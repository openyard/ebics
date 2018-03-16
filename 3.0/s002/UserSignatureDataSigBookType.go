// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
