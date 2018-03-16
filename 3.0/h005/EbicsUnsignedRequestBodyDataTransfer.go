// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type EbicsUnsignedRequestBodyDataTransfer struct {
	SignatureData *EbicsUnsignedRequestBodyDataTransferSignatureData `xml:"SignatureData"`
	OrderData     *EbicsUnsignedRequestBodyDataTransferOrderData     `xml:"OrderData"`
}

func NewEbicsUnsignedRequestBodyDataTransfer() *EbicsUnsignedRequestBodyDataTransfer {
	return new(EbicsUnsignedRequestBodyDataTransfer)
}

func (me *EbicsUnsignedRequestBodyDataTransfer) SetSignatureData(value *EbicsUnsignedRequestBodyDataTransferSignatureData) {
	me.SignatureData = value
}

func (me *EbicsUnsignedRequestBodyDataTransfer) AddSignatureData() *EbicsUnsignedRequestBodyDataTransferSignatureData {
	me.SignatureData = new(EbicsUnsignedRequestBodyDataTransferSignatureData)
	return me.SignatureData
}

func (me *EbicsUnsignedRequestBodyDataTransfer) SetOrderData(value *EbicsUnsignedRequestBodyDataTransferOrderData) {
	me.OrderData = value
}

func (me *EbicsUnsignedRequestBodyDataTransfer) AddOrderData() *EbicsUnsignedRequestBodyDataTransferOrderData {
	me.OrderData = new(EbicsUnsignedRequestBodyDataTransferOrderData)
	return me.OrderData
}
