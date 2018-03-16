// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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
