// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
