// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SignaturePropertyType struct {
	Target w3c.AnyURI `xml:"Target,attr"`
	Id     w3c.ID     `xml:"Id,attr,omitempty"`

	Any []*w3c.Any
}
