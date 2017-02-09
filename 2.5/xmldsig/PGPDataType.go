// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type PGPDataType struct {
	PGPKeyID     w3c.Base64Binary `xml:"PGPKeyID"`
	PGPKeyPacket w3c.Base64Binary `xml:"PGPKeyPacket,omitempty"`

	Any []*w3c.Any
}
