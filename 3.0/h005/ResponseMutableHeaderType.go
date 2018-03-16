// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type ResponseMutableHeaderType struct {
	TransactionPhase *TransactionPhaseType                   `xml:"TransactionPhase"`
	SegmentNumber    *ResponseMutableHeaderTypeSegmentNumber `xml:"SegmentNumber,omitempty"`
	OrderID          *OrderIDType                            `xml:"OrderID,omitempty"`
	ReturnCode       *ReturnCodeType                         `xml:"ReturnCode"`
	ReportText       *ReportTextType                         `xml:"ReportText"`

	Any []*w3c.Any
}

func NewResponseMutableHeaderType() *ResponseMutableHeaderType {
	return new(ResponseMutableHeaderType)
}

func (me *ResponseMutableHeaderType) SetTransactionPhase(value *TransactionPhaseType) {
	me.TransactionPhase = value
}

func (me *ResponseMutableHeaderType) AddTransactionPhase() *TransactionPhaseType {
	me.TransactionPhase = new(TransactionPhaseType)
	return me.TransactionPhase
}

func (me *ResponseMutableHeaderType) SetSegmentNumber(value *ResponseMutableHeaderTypeSegmentNumber) {
	me.SegmentNumber = value
}

func (me *ResponseMutableHeaderType) AddSegmentNumber() *ResponseMutableHeaderTypeSegmentNumber {
	me.SegmentNumber = new(ResponseMutableHeaderTypeSegmentNumber)
	return me.SegmentNumber
}

func (me *ResponseMutableHeaderType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *ResponseMutableHeaderType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *ResponseMutableHeaderType) SetReturnCode(value *ReturnCodeType) {
	me.ReturnCode = value
}

func (me *ResponseMutableHeaderType) AddReturnCode() *ReturnCodeType {
	me.ReturnCode = new(ReturnCodeType)
	return me.ReturnCode
}

func (me *ResponseMutableHeaderType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}

func (me *ResponseMutableHeaderType) AddReportText() *ReportTextType {
	me.ReportText = new(ReportTextType)
	return me.ReportText
}
