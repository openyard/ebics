// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
