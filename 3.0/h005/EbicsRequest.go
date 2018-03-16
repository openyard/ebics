// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *EbicsRequest) SetBody(value *EbicsRequestBody) {
	me.Body = value
}

func (me *EbicsRequest) SetAuthSignature(value *AuthSignature) {
	me.AuthSignature = value
}
