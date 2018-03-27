// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

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
