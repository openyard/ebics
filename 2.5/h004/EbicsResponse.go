// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type EbicsResponse struct {
	Header        *EbicsResponseHeader `xml:"header"`
	Body          *EbicsResponseBody   `xml:"body"`
	AuthSignature *AuthSignature       `xml:"AuthSignature"`
	VersionAttrGroup
}

func NewEbicsResponse() *EbicsResponse {
	return new(EbicsResponse)
}

func (me *EbicsResponse) SetHeader(value *EbicsResponseHeader) {
	me.Header = value
}

func (me *EbicsResponse) AddHeader() *EbicsResponseHeader {
	me.Header = new(EbicsResponseHeader)
	return me.Header
}

func (me *EbicsResponse) SetBody(value *EbicsResponseBody) {
	me.Body = value
}

func (me *EbicsResponse) AddBody() *EbicsResponseBody {
	me.Body = new(EbicsResponseBody)
	return me.Body
}

func (me *EbicsResponse) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
