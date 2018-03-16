// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type EbicsUnsignedRequestBody struct {
	DataTransfer *EbicsUnsignedRequestBodyDataTransfer `xml:"DataTransfer"`
}

func NewEbicsUnsignedRequestBody() *EbicsUnsignedRequestBody {
	return new(EbicsUnsignedRequestBody)
}

func (me *EbicsUnsignedRequestBody) SetDataTransfer(value *EbicsUnsignedRequestBodyDataTransfer) {
	me.DataTransfer = value
}
