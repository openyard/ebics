// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVTOrderInfoType struct {
	MsgName       *MessageType                   `xml:"MsgName,omitempty"`
	AccountInfo   *HVTAccountInfoType            `xml:"AccountInfo"`
	ExecutionDate *HVTOrderInfoTypeExecutionDate `xml:"ExecutionDate,omitempty"`
	Amount        *HVTOrderInfoTypeAmount        `xml:"Amount"`
	Description   *HVTOrderInfoTypeDescription   `xml:"Description,omitempty"`

	Any []*w3c.Any
}

func NewHVTOrderInfoType() *HVTOrderInfoType {
	return new(HVTOrderInfoType)
}

func (me *HVTOrderInfoType) SetMsgName(value *MessageType) {
	me.MsgName = value
}

func (me *HVTOrderInfoType) SetAccountInfo(value *HVTAccountInfoType) {
	me.AccountInfo = value
}

func (me *HVTOrderInfoType) SetExecutionDate(value *HVTOrderInfoTypeExecutionDate) {
	me.ExecutionDate = value
}

func (me *HVTOrderInfoType) SetAmount(value *HVTOrderInfoTypeAmount) {
	me.Amount = value
}

func (me *HVTOrderInfoType) SetDescription(value *HVTOrderInfoTypeDescription) {
	me.Description = value
}
