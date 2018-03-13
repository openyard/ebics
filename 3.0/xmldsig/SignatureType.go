// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SignatureType struct {
	Id             w3c.ID         `xml:"Id,attr,omitempty"`
	SignedInfo     SignedInfo     `xml:"SignedInfo"`
	SignatureValue SignatureValue `xml:"SignatureValue"`
	KeyInfo        KeyInfo        `xml:"KeyInfo,omitempty"`
	Object         Object         `xml:"Object,omitempty"`
}
