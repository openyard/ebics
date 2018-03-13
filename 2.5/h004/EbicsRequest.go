// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type EbicsRequest struct {
    
    Header EbicsRequestHeader `xml:"header"`
    Body EbicsRequestBody `xml:"body"`
    AuthSignature AuthSignature `xml:"AuthSignature"`
    VersionAttrGroup
    }
