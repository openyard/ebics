// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type EbicsUnsecuredRequest struct {
	Header *EbicsUnsecuredRequestHeader `xml:"header"`
	Body   *EbicsUnsecuredRequestBody   `xml:"body"`
	VersionAttrGroup
}

func NewEbicsUnsecuredRequest() *EbicsUnsecuredRequest {
	return new(EbicsUnsecuredRequest)
}

func (me *EbicsUnsecuredRequest) SetHeader(value *EbicsUnsecuredRequestHeader) {
	me.Header = value
}

func (me *EbicsUnsecuredRequest) AddHeader() *EbicsUnsecuredRequestHeader {
	me.Header = new(EbicsUnsecuredRequestHeader)
	return me.Header
}

func (me *EbicsUnsecuredRequest) SetBody(value *EbicsUnsecuredRequestBody) {
	me.Body = value
}

func (me *EbicsUnsecuredRequest) AddBody() *EbicsUnsecuredRequestBody {
	me.Body = new(EbicsUnsecuredRequestBody)
	return me.Body
}
