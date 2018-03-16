// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type EbicsKeyManagementResponse struct {
	Header *EbicsKeyManagementResponseHeader `xml:"header"`
	Body   *EbicsKeyManagementResponseBody   `xml:"body"`
	VersionAttrGroup
}

func NewEbicsKeyManagementResponse() *EbicsKeyManagementResponse {
	return new(EbicsKeyManagementResponse)
}

func (me *EbicsKeyManagementResponse) SetHeader(value *EbicsKeyManagementResponseHeader) {
	me.Header = value
}

func (me *EbicsKeyManagementResponse) SetBody(value *EbicsKeyManagementResponseBody) {
	me.Body = value
}
