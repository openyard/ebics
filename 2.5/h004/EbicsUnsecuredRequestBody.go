// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type EbicsUnsecuredRequestBody struct {
	DataTransfer *EbicsUnsecuredRequestBodyDataTransfer `xml:"DataTransfer"`
}

func NewEbicsUnsecuredRequestBody() *EbicsUnsecuredRequestBody {
	return new(EbicsUnsecuredRequestBody)
}

func (me *EbicsUnsecuredRequestBody) SetDataTransfer(value *EbicsUnsecuredRequestBodyDataTransfer) {
	me.DataTransfer = value
}

func (me *EbicsUnsecuredRequestBody) AddDataTransfer() *EbicsUnsecuredRequestBodyDataTransfer {
	me.DataTransfer = new(EbicsUnsecuredRequestBodyDataTransfer)
	return me.DataTransfer
}
