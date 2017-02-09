// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type EbicsUnsignedRequest struct {
	Header EbicsUnsignedRequestHeader `xml:"header"`
	Body   EbicsUnsignedRequestBody   `xml:"body"`
	VersionAttrGroup
}
