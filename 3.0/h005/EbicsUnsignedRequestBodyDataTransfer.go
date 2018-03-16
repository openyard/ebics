// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *EbicsUnsignedRequestBodyDataTransfer) SetOrderData(value *EbicsUnsignedRequestBodyDataTransferOrderData) {
	me.OrderData = value
}
