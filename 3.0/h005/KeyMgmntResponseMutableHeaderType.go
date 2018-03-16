// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type KeyMgmntResponseMutableHeaderType struct {
	OrderID    *OrderIDType    `xml:"OrderID,omitempty"`
	ReturnCode *ReturnCodeType `xml:"ReturnCode"`
	ReportText *ReportTextType `xml:"ReportText"`

	Any []*w3c.Any
}

func NewKeyMgmntResponseMutableHeaderType() *KeyMgmntResponseMutableHeaderType {
	return new(KeyMgmntResponseMutableHeaderType)
}

func (me *KeyMgmntResponseMutableHeaderType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *KeyMgmntResponseMutableHeaderType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *KeyMgmntResponseMutableHeaderType) SetReturnCode(value *ReturnCodeType) {
	me.ReturnCode = value
}

func (me *KeyMgmntResponseMutableHeaderType) AddReturnCode() *ReturnCodeType {
	me.ReturnCode = new(ReturnCodeType)
	return me.ReturnCode
}

func (me *KeyMgmntResponseMutableHeaderType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}

func (me *KeyMgmntResponseMutableHeaderType) AddReportText() *ReportTextType {
	me.ReportText = new(ReportTextType)
	return me.ReportText
}
