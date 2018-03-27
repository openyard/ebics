// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
