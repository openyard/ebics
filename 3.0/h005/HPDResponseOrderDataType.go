// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type HPDResponseOrderDataType struct {
	AccessParams   *HPDAccessParamsType   `xml:"AccessParams"`
	ProtocolParams *HPDProtocolParamsType `xml:"ProtocolParams"`
}

func NewHPDResponseOrderDataType() *HPDResponseOrderDataType {
	return new(HPDResponseOrderDataType)
}

func (me *HPDResponseOrderDataType) SetAccessParams(value *HPDAccessParamsType) {
	me.AccessParams = value
}

func (me *HPDResponseOrderDataType) SetProtocolParams(value *HPDProtocolParamsType) {
	me.ProtocolParams = value
}
