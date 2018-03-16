// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsUnsignedRequest struct {
	Header *EbicsUnsignedRequestHeader `xml:"header"`
	Body   *EbicsUnsignedRequestBody   `xml:"body"`
	VersionAttrGroup
}

func NewEbicsUnsignedRequest() *EbicsUnsignedRequest {
	return new(EbicsUnsignedRequest)
}

func (me *EbicsUnsignedRequest) SetHeader(value *EbicsUnsignedRequestHeader) {
	me.Header = value
}

func (me *EbicsUnsignedRequest) SetBody(value *EbicsUnsignedRequestBody) {
	me.Body = value
}
