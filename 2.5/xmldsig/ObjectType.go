// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ObjectType struct {
	Id       w3c.ID     `xml:"Id,attr,omitempty"`
	MimeType w3c.String `xml:"MimeType,attr,omitempty"`
	Encoding w3c.AnyURI `xml:"Encoding,attr,omitempty"`

	Any []*w3c.Any
}
