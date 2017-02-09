// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type EbicsRequest struct {
	Header        EbicsRequestHeader `xml:"header"`
	Body          EbicsRequestBody   `xml:"body"`
	AuthSignature AuthSignature      `xml:"AuthSignature"`
	VersionAttrGroup
}
