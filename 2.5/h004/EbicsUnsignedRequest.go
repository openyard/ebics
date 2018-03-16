// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexType
type EbicsUnsignedRequest struct {
	Header *EbicsUnsignedRequestHeader `xml:"header"`
	Body   *EbicsUnsignedRequestBody   `xml:"body"`
	VersionAttrGroup
}

func NewEbicsUnsignedRequest() *EbicsUnsignedRequest {
	return new(EbicsUnsignedRequest)
}

func (me *EbicsUnsignedRequest) SetHeader(value *EbicsUnsignedRequestHeader) {
	me.Header = value
}

func (me *EbicsUnsignedRequest) AddHeader() *EbicsUnsignedRequestHeader {
	me.Header = new(EbicsUnsignedRequestHeader)
	return me.Header
}

func (me *EbicsUnsignedRequest) SetBody(value *EbicsUnsignedRequestBody) {
	me.Body = value
}

func (me *EbicsUnsignedRequest) AddBody() *EbicsUnsignedRequestBody {
	me.Body = new(EbicsUnsignedRequestBody)
	return me.Body
}
