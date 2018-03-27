// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *HVTOrderInfoType) AddMsgName() *MessageType {
	me.MsgName = new(MessageType)
	return me.MsgName
}

func (me *HVTOrderInfoType) SetAccountInfo(value *HVTAccountInfoType) {
	me.AccountInfo = value
}

func (me *HVTOrderInfoType) AddAccountInfo() *HVTAccountInfoType {
	me.AccountInfo = new(HVTAccountInfoType)
	return me.AccountInfo
}

func (me *HVTOrderInfoType) SetExecutionDate(value *HVTOrderInfoTypeExecutionDate) {
	me.ExecutionDate = value
}

func (me *HVTOrderInfoType) AddExecutionDate() *HVTOrderInfoTypeExecutionDate {
	me.ExecutionDate = new(HVTOrderInfoTypeExecutionDate)
	return me.ExecutionDate
}

func (me *HVTOrderInfoType) SetAmount(value *HVTOrderInfoTypeAmount) {
	me.Amount = value
}

func (me *HVTOrderInfoType) AddAmount() *HVTOrderInfoTypeAmount {
	me.Amount = new(HVTOrderInfoTypeAmount)
	return me.Amount
}

func (me *HVTOrderInfoType) SetDescription(value *HVTOrderInfoTypeDescription) {
	me.Description = value
}

func (me *HVTOrderInfoType) AddDescription() *HVTOrderInfoTypeDescription {
	me.Description = new(HVTOrderInfoTypeDescription)
	return me.Description
}
