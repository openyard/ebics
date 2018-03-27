// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type EbicsUnsecuredRequestBodyDataTransfer struct {
	OrderData *EbicsUnsecuredRequestBodyDataTransferOrderData `xml:"OrderData"`
}

func NewEbicsUnsecuredRequestBodyDataTransfer() *EbicsUnsecuredRequestBodyDataTransfer {
	return new(EbicsUnsecuredRequestBodyDataTransfer)
}

func (me *EbicsUnsecuredRequestBodyDataTransfer) SetOrderData(value *EbicsUnsecuredRequestBodyDataTransferOrderData) {
	me.OrderData = value
}

func (me *EbicsUnsecuredRequestBodyDataTransfer) AddOrderData() *EbicsUnsecuredRequestBodyDataTransferOrderData {
	me.OrderData = new(EbicsUnsecuredRequestBodyDataTransferOrderData)
	return me.OrderData
}
