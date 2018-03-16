// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
