// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVTResponseOrderDataType struct {
	NumOrderInfos *NumOrderInfosType  `xml:"NumOrderInfos"`
	OrderInfo     []*HVTOrderInfoType `xml:"OrderInfo"`

	Any []*w3c.Any
}

func NewHVTResponseOrderDataType() *HVTResponseOrderDataType {
	return new(HVTResponseOrderDataType)
}

func (me *HVTResponseOrderDataType) SetNumOrderInfos(value *NumOrderInfosType) {
	me.NumOrderInfos = value
}

func (me *HVTResponseOrderDataType) AddNumOrderInfos() *NumOrderInfosType {
	me.NumOrderInfos = new(NumOrderInfosType)
	return me.NumOrderInfos
}

func (me *HVTResponseOrderDataType) AddOrderInfo(value *HVTOrderInfoType) {
	me.OrderInfo = append(me.OrderInfo, value)
}
