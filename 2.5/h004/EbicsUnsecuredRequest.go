// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

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

func (me *EbicsUnsecuredRequest) SetBody(value *EbicsUnsecuredRequestBody) {
	me.Body = value
}
