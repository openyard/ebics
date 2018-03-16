// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *EbicsKeyManagementResponse) AddHeader() *EbicsKeyManagementResponseHeader {
	me.Header = new(EbicsKeyManagementResponseHeader)
	return me.Header
}

func (me *EbicsKeyManagementResponse) SetBody(value *EbicsKeyManagementResponseBody) {
	me.Body = value
}

func (me *EbicsKeyManagementResponse) AddBody() *EbicsKeyManagementResponseBody {
	me.Body = new(EbicsKeyManagementResponseBody)
	return me.Body
}
