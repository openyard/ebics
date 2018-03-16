// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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
