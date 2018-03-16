// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *TransferReceiptResponseType) SetTimestampBankParameter(value *TimestampType) {
	me.TimestampBankParameter = value
}
