// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

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

func (me *EbicsNoPubKeyDigestsRequest) SetBody(value *EbicsNoPubKeyDigestsRequestBody) {
	me.Body = value
}

func (me *EbicsNoPubKeyDigestsRequest) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
