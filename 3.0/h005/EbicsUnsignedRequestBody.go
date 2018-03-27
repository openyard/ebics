// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

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

func (me *EbicsUnsignedRequestBody) AddDataTransfer() *EbicsUnsignedRequestBodyDataTransfer {
	me.DataTransfer = new(EbicsUnsignedRequestBodyDataTransfer)
	return me.DataTransfer
}
