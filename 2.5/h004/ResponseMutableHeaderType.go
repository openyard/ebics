// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *ResponseMutableHeaderType) SetSegmentNumber(value *ResponseMutableHeaderTypeSegmentNumber) {
	me.SegmentNumber = value
}

func (me *ResponseMutableHeaderType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *ResponseMutableHeaderType) SetReturnCode(value *ReturnCodeType) {
	me.ReturnCode = value
}

func (me *ResponseMutableHeaderType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}
