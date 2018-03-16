// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexElement
type EbicsRequestBodyTransferReceipt struct {
	Value *TransferReceiptRequestType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsRequestBodyTransferReceipt() *EbicsRequestBodyTransferReceipt {
	return new(EbicsRequestBodyTransferReceipt)
}
