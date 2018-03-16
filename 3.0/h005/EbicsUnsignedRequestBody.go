// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
