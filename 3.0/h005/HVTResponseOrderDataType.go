// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HVTResponseOrderDataType) AddOrderInfo(value *HVTOrderInfoType) {
	me.OrderInfo = append(me.OrderInfo, value)
}
