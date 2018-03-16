// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
