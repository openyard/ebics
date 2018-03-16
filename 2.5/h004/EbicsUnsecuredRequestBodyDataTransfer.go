// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
