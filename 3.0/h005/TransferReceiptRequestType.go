// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type TransferReceiptRequestType struct {
	ReceiptCode *ReceiptCodeType `xml:"ReceiptCode"`

	Any []*w3c.Any
}

func NewTransferReceiptRequestType() *TransferReceiptRequestType {
	return new(TransferReceiptRequestType)
}

func (me *TransferReceiptRequestType) SetReceiptCode(value *ReceiptCodeType) {
	me.ReceiptCode = value
}

func (me *TransferReceiptRequestType) AddReceiptCode() *ReceiptCodeType {
	me.ReceiptCode = new(ReceiptCodeType)
	return me.ReceiptCode
}
