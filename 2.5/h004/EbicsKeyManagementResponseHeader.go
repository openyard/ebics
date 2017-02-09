// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type EbicsKeyManagementResponseHeader struct {
	Static  EbicsKeyManagementResponseHeaderStatic `xml:"static"`
	Mutable KeyMgmntResponseMutableHeaderType      `xml:"mutable"`
	AuthenticationMarker
}
