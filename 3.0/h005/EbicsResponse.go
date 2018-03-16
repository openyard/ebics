// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

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

func (me *EbicsResponse) SetBody(value *EbicsResponseBody) {
	me.Body = value
}

func (me *EbicsResponse) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
