// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
