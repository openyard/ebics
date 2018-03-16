// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type TransferReceiptResponseType struct {
	ReturnCodeReceipt      *ReturnCodeType `xml:"ReturnCodeReceipt"`
	TimestampBankParameter *TimestampType  `xml:"TimestampBankParameter,omitempty"`

	Any []*w3c.Any
}

func NewTransferReceiptResponseType() *TransferReceiptResponseType {
	return new(TransferReceiptResponseType)
}

func (me *TransferReceiptResponseType) SetReturnCodeReceipt(value *ReturnCodeType) {
	me.ReturnCodeReceipt = value
}

func (me *TransferReceiptResponseType) AddReturnCodeReceipt() *ReturnCodeType {
	me.ReturnCodeReceipt = new(ReturnCodeType)
	return me.ReturnCodeReceipt
}

func (me *TransferReceiptResponseType) SetTimestampBankParameter(value *TimestampType) {
	me.TimestampBankParameter = value
}

func (me *TransferReceiptResponseType) AddTimestampBankParameter() *TimestampType {
	me.TimestampBankParameter = new(TimestampType)
	return me.TimestampBankParameter
}
