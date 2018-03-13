// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type CanonicalizationMethodType struct {
	Algorithm w3c.AnyURI `xml:"Algorithm,attr"`

	Any []*w3c.Any
}
