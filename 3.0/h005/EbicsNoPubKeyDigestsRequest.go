// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type EbicsNoPubKeyDigestsRequest struct {
	Header        *EbicsNoPubKeyDigestsRequestHeader `xml:"header"`
	Body          *EbicsNoPubKeyDigestsRequestBody   `xml:"body"`
	AuthSignature *AuthSignature                     `xml:"AuthSignature"`
	VersionAttrGroup
}

func NewEbicsNoPubKeyDigestsRequest() *EbicsNoPubKeyDigestsRequest {
	return new(EbicsNoPubKeyDigestsRequest)
}

func (me *EbicsNoPubKeyDigestsRequest) SetHeader(value *EbicsNoPubKeyDigestsRequestHeader) {
	me.Header = value
}

func (me *EbicsNoPubKeyDigestsRequest) AddHeader() *EbicsNoPubKeyDigestsRequestHeader {
	me.Header = new(EbicsNoPubKeyDigestsRequestHeader)
	return me.Header
}

func (me *EbicsNoPubKeyDigestsRequest) SetBody(value *EbicsNoPubKeyDigestsRequestBody) {
	me.Body = value
}

func (me *EbicsNoPubKeyDigestsRequest) AddBody() *EbicsNoPubKeyDigestsRequestBody {
	me.Body = new(EbicsNoPubKeyDigestsRequestBody)
	return me.Body
}

func (me *EbicsNoPubKeyDigestsRequest) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
