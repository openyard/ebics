// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type EbicsRequest struct {
	Header        *EbicsRequestHeader `xml:"header"`
	Body          *EbicsRequestBody   `xml:"body"`
	AuthSignature *AuthSignature      `xml:"AuthSignature"`
	VersionAttrGroup
}

func NewEbicsRequest() *EbicsRequest {
	return new(EbicsRequest)
}

func (me *EbicsRequest) SetHeader(value *EbicsRequestHeader) {
	me.Header = value
}

func (me *EbicsRequest) AddHeader() *EbicsRequestHeader {
	me.Header = new(EbicsRequestHeader)
	return me.Header
}

func (me *EbicsRequest) SetBody(value *EbicsRequestBody) {
	me.Body = value
}

func (me *EbicsRequest) AddBody() *EbicsRequestBody {
	me.Body = new(EbicsRequestBody)
	return me.Body
}

func (me *EbicsRequest) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
