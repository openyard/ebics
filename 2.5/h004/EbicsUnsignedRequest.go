// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
