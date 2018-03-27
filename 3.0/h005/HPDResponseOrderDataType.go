// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *HPDResponseOrderDataType) AddAccessParams() *HPDAccessParamsType {
	me.AccessParams = new(HPDAccessParamsType)
	return me.AccessParams
}

func (me *HPDResponseOrderDataType) SetProtocolParams(value *HPDProtocolParamsType) {
	me.ProtocolParams = value
}

func (me *HPDResponseOrderDataType) AddProtocolParams() *HPDProtocolParamsType {
	me.ProtocolParams = new(HPDProtocolParamsType)
	return me.ProtocolParams
}
