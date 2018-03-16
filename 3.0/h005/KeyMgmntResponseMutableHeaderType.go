// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *KeyMgmntResponseMutableHeaderType) SetReturnCode(value *ReturnCodeType) {
	me.ReturnCode = value
}

func (me *KeyMgmntResponseMutableHeaderType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}
