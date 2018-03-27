// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
