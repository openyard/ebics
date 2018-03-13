// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type PGPDataType struct {
	PGPKeyID     w3c.Base64Binary `xml:"PGPKeyID"`
	PGPKeyPacket w3c.Base64Binary `xml:"PGPKeyPacket,omitempty"`

	Any []*w3c.Any
}
